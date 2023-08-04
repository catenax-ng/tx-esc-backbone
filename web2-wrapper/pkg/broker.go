package web2wrapper

import (
	"context"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	"io"
)

type MOD string

const (
	CREATE MOD = "create"
	UPDATE MOD = "update"
	DELETE MOD = "delete"
)

type SENDER string

const (
	WRAPPER SENDER = "wrapper"
	CLIENT  SENDER = "client"
)

type Msg struct {
	Res types.Resource `json:"res"`
	Mod MOD            `json:"mod"`
	Src SENDER         `json:"src"`
}

type Broker interface {
	io.Closer
	Submit(msg *Msg) error
	Receive(ctx context.Context) <-chan *Msg
}
