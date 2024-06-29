package keeper

import (
	"certifichain/x/issuer/types"
)

var _ types.QueryServer = Keeper{}
