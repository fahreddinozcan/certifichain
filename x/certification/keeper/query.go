package keeper

import (
	"certifichain/x/certification/types"
)

var _ types.QueryServer = Keeper{}
