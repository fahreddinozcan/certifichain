package learner_test

import (
	"testing"

	keepertest "certifichain/testutil/keeper"
	"certifichain/testutil/nullify"
	learner "certifichain/x/learner/module"
	"certifichain/x/learner/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LearnerKeeper(t)
	learner.InitGenesis(ctx, k, genesisState)
	got := learner.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
