package web2wrapper

import (
	"context"
	"cosmossdk.io/log"
	"fmt"
)

type ducttape struct {
	queue2chain queue2chain
	chain2queue chain2queue
}

type Ducttape interface {
	Forward(ctx context.Context, startBlock int64)
}

type queue2chain struct {
	logger log.Logger
	src    Broker
	trg    ResourceSyncClient
}

type chain2queue struct {
	logger log.Logger
	src    ResourceSyncClient
	trg    Broker
}

func NewDucttape(broker Broker, client ResourceSyncClient, logger log.Logger) Ducttape {
	return &ducttape{
		queue2chain: queue2chain{
			logger: logger,
			src:    broker,
			trg:    client,
		},
		chain2queue: chain2queue{
			logger: logger,
			src:    client,
			trg:    broker,
		},
	}
}

func (d *ducttape) Forward(ctx context.Context, startBlock int64) {
	go d.chain2queue.forward(ctx, startBlock)
	go d.queue2chain.forward(ctx)
	select {
	case <-ctx.Done():
	}
}

func (c *chain2queue) forward(ctx context.Context, startBlock int64) {
	for msg := range c.src.Poll(ctx, startBlock) {
		msg.Src = WRAPPER
		err := c.trg.Submit(msg)
		if err != nil {
			c.logger.Error(fmt.Sprintf("failed submitting %s - reason %v", msg, err))
		}
	}
}

func (q *queue2chain) forward(ctx context.Context) {
	for msg := range q.src.Receive(ctx) {
		if msg.Src == WRAPPER {
			q.logger.Debug(fmt.Sprintf("Ignore self submitted msg %s", msg))
			continue
		}
		switch msg.Mod {
		case CREATE:
			res, err := q.trg.CreateResource(ctx, msg.Res)
			if err != nil {
				q.logger.Error(fmt.Sprintf("Failed to create resource %s -  %v", msg, err))
			} else {
				q.logger.Debug(fmt.Sprintf("Created resource: %s", res))
			}
		case UPDATE:
			res, err := q.trg.UpdateResource(ctx, msg.Res)
			if err != nil {
				q.logger.Error(fmt.Sprintf("Failed to update resource %s -  %v", msg, err))
			} else {
				q.logger.Debug(fmt.Sprintf("Updated resource: %s", res))
			}
		case DELETE:
			res, err := q.trg.DeleteResource(ctx, msg.Res.OrigResId)
			if err != nil {
				q.logger.Error(fmt.Sprintf("Failed to delete resource %s -  %v", msg, err))
			} else {
				q.logger.Debug(fmt.Sprintf("Deleted resource: %s", res))
			}
		default:
			q.logger.Debug(fmt.Sprintf("Unknown modification %s on %s", msg.Mod, msg))
		}
	}
}
