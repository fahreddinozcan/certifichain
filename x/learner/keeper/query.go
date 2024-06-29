package keeper

import (
	"certifichain/x/learner/types"
)

var _ types.QueryServer = Keeper{}
