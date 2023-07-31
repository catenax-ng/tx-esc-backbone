// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package web2wrapper

import (
	"context"
	"encoding/base64"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"log"
	"regexp"
)

var resourceEventTypeRegex = regexp.MustCompile(`^catenax\.escbackbone\.resourcesync\.Event.+`)

// taken from ignite generated code - end

type ResourceSyncClient interface {
	CreateResource(ctx context.Context, resource types.Resource) (cosmosclient.Response, error)
	UpdateResource(ctx context.Context, resource types.Resource) (cosmosclient.Response, error)
	DeleteResource(ctx context.Context, origResId string) (cosmosclient.Response, error)
	QueryAllResources(ctx context.Context) (*types.QueryAllResourceMapResponse, error)
	Poll(ctx context.Context)
}

type chainClient struct {
	client        cosmosclient.Client
	from          cosmosaccount.Account
	addressPrefix string
	queryClient   types.QueryClient
	blockHeight   int64
}

func NewChainClient(ctx context.Context, config Config) (ResourceSyncClient, error) {
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
		log.Fatal(err)
	}
	account, err := client.Account(config.From)
	if err != nil {
		log.Fatal(err)
	}
	return &chainClient{
		client:        client,
		from:          account,
		addressPrefix: config.AddressPrefix,
		queryClient:   types.NewQueryClient(client.Context()),
		blockHeight:   config.StartBlock,
	}, nil
}

func (c *chainClient) mustGetAddress() (addr string) {
	addr, err := c.from.Address(c.addressPrefix)
	if err != nil {
		log.Fatal(err)
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

func (c *chainClient) Poll(ctx context.Context) {
	height, err := c.client.LatestBlockHeight(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if c.blockHeight == height {
		log.Printf("uptodate")
		return
	}
	if c.blockHeight < height-100 {
		log.Printf("Way behind: %d -> %d", c.blockHeight, height)
	}
	defer c.SafeConfig()
	for c.blockHeight < height {
		blockProcessed := c.blockHeight + 1
		txs, err := c.client.GetBlockTXs(ctx, blockProcessed)
		if err != nil {
			log.Fatalf("Cannot fetch block %d", blockProcessed)
		}
		log.Printf("Processing block %d\r", blockProcessed)
		c.parseBlock(txs)
		c.blockHeight = blockProcessed
	}
}

func (c *chainClient) parseBlock(txs []cosmosclient.TX) {

	for _, tx := range txs {
		resourceEvents := make([]abci.Event, 0)
		for _, event := range tx.Raw.TxResult.Events {
			if resourceEventTypeRegex.MatchString(event.Type) {
				resourceEvents = append(resourceEvents, event)
			}
		}
		if len(resourceEvents) > 0 {
			log.Printf("Block: %d, Tx: %s, Events:", tx.Raw.Height, tx.Raw.Hash)

		}
		for _, resourceEvent := range resourceEvents {
			oldAttributes := resourceEvent.Attributes
			reworkedAttributes := make([]abci.EventAttribute, 0)
			for _, a := range oldAttributes {
				reworkedAttributes = append(reworkedAttributes, abci.EventAttribute{
					Key:   string(mustDecodeBase64(a.Key)),
					Value: string(mustDecodeBase64(a.Value)),
					Index: a.Index,
				})
			}
			resourceEvent.Attributes = reworkedAttributes
			event, err := sdk.ParseTypedEvent(resourceEvent)

			if err != nil {
				log.Fatal(err)
			}
			log.Println(event)
		}
	}
}

func mustDecodeBase64(str string) []byte {
	result, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
