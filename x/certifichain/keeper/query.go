package keeper

import (
	"certifichain/x/certifichain/types"
)

var _ types.QueryServer = Keeper{}
