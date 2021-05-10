// DONTCOVER
package v036

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const ModuleName = "supply"

type (
	GenesisState struct {
		Supply sdk.Coins `json:"supply" yaml:"supply"`
	}
)
