// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package web2wrapper

import (
	"context"
	"cosmossdk.io/log"
	"encoding/base64"
	"fmt"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/coreos/go-semver/semver"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_version "github.com/cosmos/cosmos-sdk/version"
	proto "github.com/cosmos/gogoproto/proto"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"regexp"
	"syscall"
	"time"
)

func cosmosSDKatLeastv047() bool {
	vSdkVer := sdk_version.NewInfo().CosmosSdkVersion
	//cut of the v
	sdkVer := vSdkVer[1:]
	current := semver.New(sdkVer)
	v047 := semver.New("0.47.0")
	return !current.LessThan(*v047)
}

var skdVersionAtLeast047 = cosmosSDKatLeastv047()

func createEventTypeRegex() *regexp.Regexp {
	if skdVersionAtLeast047 {
		return regexp.MustCompile(`^escbackbone\.resourcesync\.Event.+`)
	} else {
		return regexp.MustCompile(`^catenax\.escbackbone\.resourcesync\.Event.+`)
	}
}

var resourceEventTypeRegex = createEventTypeRegex()

type ResourceSyncClient interface {
	CreateResource(ctx context.Context, resource types.Resource) (cosmosclient.Response, error)
	UpdateResource(ctx context.Context, resource types.Resource) (cosmosclient.Response, error)
	DeleteResource(ctx context.Context, origResId string) (cosmosclient.Response, error)
	QueryAllResources(ctx context.Context) (*types.QueryAllResourceMapResponse, error)
	Poll(ctx context.Context, startBlock int64) <-chan *Msg
}

type chainClient struct {
	client        cosmosclient.Client
	from          cosmosaccount.Account
	addressPrefix string
	queryClient   types.QueryClient
	blockHeight   int64
	logger        log.Logger
}

func NewChainClient(ctx context.Context, logger log.Logger, config *Config) (ResourceSyncClient, error) {
	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx,
		cosmosclient.WithAddressPrefix(config.AddressPrefix),
		cosmosclient.WithNodeAddress(config.NodeAddress),
		cosmosclient.WithFees(config.Fees),
		cosmosclient.WithGas(config.Gas),
		cosmosclient.WithHome(config.Home),
		cosmosclient.WithKeyringBackend(config.KeyRingBackend),
	)
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot create cosmos client - %v", err))
		return nil, err
	}
	account, err := client.Account(config.From)
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot read account from config - %v", err))
		return nil, err
	}
	return &chainClient{
		client:        client,
		from:          account,
		addressPrefix: config.AddressPrefix,
		queryClient:   types.NewQueryClient(client.Context()),
		blockHeight:   config.StartBlock,
		logger:        logger,
	}, nil
}

func (c *chainClient) mustGetAddress() (addr string) {
	addr, err := c.from.Address(c.addressPrefix)
	if err != nil {
		panic("Cannot access address for account")
	}
	return addr
}

func (c *chainClient) CreateResource(ctx context.Context, resource types.Resource) (cosmosclient.Response, error) {
	resource.Originator = c.mustGetAddress()
	msg := types.NewMsgCreateResource(
		c.mustGetAddress(),
		&resource,
	)
	return c.client.BroadcastTx(ctx, c.from, msg)
}

func (c *chainClient) UpdateResource(ctx context.Context, resource types.Resource) (cosmosclient.Response, error) {
	resource.Originator = c.mustGetAddress()
	msg := types.NewMsgUpdateResource(
		c.mustGetAddress(),
		&resource,
	)
	return c.client.BroadcastTx(ctx, c.from, msg)
}

func (c *chainClient) DeleteResource(ctx context.Context, origResId string) (cosmosclient.Response, error) {
	originator := c.mustGetAddress()
	msg := types.NewMsgDeleteResource(
		originator,
		originator,
		origResId,
	)
	return c.client.BroadcastTx(ctx, c.from, msg)
}

func (c *chainClient) QueryAllResources(ctx context.Context) (*types.QueryAllResourceMapResponse, error) {
	return c.queryClient.ResourceMapAll(ctx, &types.QueryAllResourceMapRequest{})
}

func (c *chainClient) Poll(ctx context.Context, startBlock int64) <-chan *Msg {
	result := make(chan *Msg)
	currentBlock := startBlock
	if currentBlock == 0 {
		currentBlock = 1
	}
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		defer close(result)
		for {
			select {
			case <-ticker.C:
				c.logger.Debug("Received tick")
				height, err := c.client.LatestBlockHeight(ctx)
				if err != nil {
					c.logger.Error(fmt.Sprintf("Cannot read current block height - %v", err))
				}
				for ; currentBlock <= height; currentBlock++ {
					//c.logger.Info(fmt.Sprintf("Processing block: %d/%d", currentBlock, height))
					c.parseBlock(ctx, currentBlock, result)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return result
}

func (c *chainClient) parseBlock(ctx context.Context, currentBlock int64, output chan<- *Msg) {
	txs, err := c.client.GetBlockTXs(ctx, currentBlock)
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to access block %d - %v", currentBlock, err))
		syscall.Exit(1)
	}
	c.logger.Debug(fmt.Sprintf("%d transactions at block %d", len(txs), currentBlock))
	for _, tx := range txs {
		resourceEvents := make([]abci.Event, 0)
		c.logger.Debug(fmt.Sprintf("%d events at transaction %d", len(tx.Raw.TxResult.Events), tx.Raw.Hash))
		for _, event := range tx.Raw.TxResult.Events {
			c.logger.Debug(fmt.Sprintf("event %s with type %s matches filter: %v", event.String(), event.Type, resourceEventTypeRegex.MatchString(event.Type)))
			if resourceEventTypeRegex.MatchString(event.Type) {
				resourceEvents = append(resourceEvents, event)
			}
		}
		if len(resourceEvents) > 0 {
			c.logger.Info(fmt.Sprintf("Block: %d, Tx: %s, Resource Events: %d", tx.Raw.Height, tx.Raw.Hash, len(resourceEvents)))
		}
		for _, resourceEvent := range resourceEvents {
			if !skdVersionAtLeast047 {
				resourceEvent.Attributes = decodeEventAttributesFromBase64(resourceEvent)
			}
			event, err := sdk.ParseTypedEvent(resourceEvent)
			if err != nil {
				c.logger.Error(fmt.Sprintf("Cannot parse and skipping:  %v", resourceEvent))
			}
			if msg := c.getResourceFrom(event); msg != nil {
				output <- msg
			}
		}
	}
}

func decodeEventAttributesFromBase64(resourceEvent abci.Event) []abci.EventAttribute {
	oldAttributes := resourceEvent.Attributes
	reworkedAttributes := make([]abci.EventAttribute, 0)
	for _, a := range oldAttributes {
		reworkedAttributes = append(reworkedAttributes, abci.EventAttribute{
			Key:   string(mustDecodeBase64(a.Key)),
			Value: string(mustDecodeBase64(a.Value)),
			Index: a.Index,
		})
	}
	return reworkedAttributes
}

func mustDecodeBase64(str string) []byte {
	result, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return result
}
func (c *chainClient) getResourceFrom(event proto.Message) *Msg {
	if created, ok := event.(*types.EventCreateResource); ok {
		return &Msg{
			Res: created.Resource,
			Mod: CREATE,
			Src: WRAPPER,
		}
	}
	if updated, ok := event.(*types.EventUpdateResource); ok {
		return &Msg{
			Res: updated.Resource,
			Mod: UPDATE,
			Src: WRAPPER,
		}
	}
	if deleted, ok := event.(*types.EventDeleteResource); ok {
		return &Msg{
			Res: deleted.Resource,
			Mod: DELETE,
			Src: WRAPPER,
		}
	}
	c.logger.Debug(fmt.Sprintf("Ignoring unknown event type %s", event))
	return nil
}
