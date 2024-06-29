package certifichain_test

import (
	"testing"

	keepertest "certifichain/testutil/keeper"
	"certifichain/testutil/nullify"
	certifichain "certifichain/x/certifichain/module"
	"certifichain/x/certifichain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CertifichainKeeper(t)
	certifichain.InitGenesis(ctx, k, genesisState)
	got := certifichain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
