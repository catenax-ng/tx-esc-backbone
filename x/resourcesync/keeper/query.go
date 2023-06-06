package keeper

import (
	"github.com/catenax/esc-backbone/x/resourcesync/types"
)

var _ types.QueryServer = Keeper{}
