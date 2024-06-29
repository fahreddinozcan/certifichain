package issuer_test

import (
	"testing"

	keepertest "certifichain/testutil/keeper"
	"certifichain/testutil/nullify"
	issuer "certifichain/x/issuer/module"
	"certifichain/x/issuer/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IssuerKeeper(t)
	issuer.InitGenesis(ctx, k, genesisState)
	got := issuer.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
