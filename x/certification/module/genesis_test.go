package certification_test

import (
	"testing"

	keepertest "certifichain/testutil/keeper"
	"certifichain/testutil/nullify"
	certification "certifichain/x/certification/module"
	"certifichain/x/certification/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CertificationKeeper(t)
	certification.InitGenesis(ctx, k, genesisState)
	got := certification.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
