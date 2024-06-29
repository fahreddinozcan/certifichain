package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "certifichain/testutil/keeper"
	"certifichain/x/learner/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LearnerKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
