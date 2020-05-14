// +build cli_test

package cli_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	cli "github.com/KiraCore/cosmos-sdk/tests/cli"
	sdk "github.com/KiraCore/cosmos-sdk/types"
	"github.com/KiraCore/cosmos-sdk/x/slashing/client/testutil"
)

func TestCLISlashingGetParams(t *testing.T) {
	t.Parallel()
	f := cli.InitFixtures(t)

	// start simd server
	proc := f.SDStart()
	t.Cleanup(func() { proc.Stop(false) })

	params := testutil.QuerySlashingParams(f)
	require.Equal(t, int64(100), params.SignedBlocksWindow)
	require.Equal(t, sdk.NewDecWithPrec(5, 1), params.MinSignedPerWindow)

	sinfo := testutil.QuerySigningInfo(f, f.SDTendermint("show-validator"))
	require.Equal(t, int64(0), sinfo.StartHeight)
	require.False(t, sinfo.Tombstoned)

	// Cleanup testing directories
	f.Cleanup()
}
