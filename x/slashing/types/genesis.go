package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	params Params, signingInfos map[string]ValidatorSigningInfo, missedBlocks map[string][]MissedBlock,
) GenesisState {

	var si = make([]SigningInfos, 0)
	for address, signingInfo := range signingInfos {
		si = append(si, SigningInfos{
			Address:      address,
			SigningInfos: signingInfo,
		})
	}

	var validatorMissedBlocks = make([]ValidatorMissedBlocks, 0)
	for address, validatorMissedBlock := range missedBlocks {
		validatorMissedBlocks = append(validatorMissedBlocks, ValidatorMissedBlocks{
			Address:      address,
			MissedBlocks: validatorMissedBlock,
		})
	}

	return GenesisState{
		Params:       params,
		SigningInfos: si,
		MissedBlocks: validatorMissedBlocks,
	}
}

// NewMissedBlock creates a new MissedBlock instance
func NewMissedBlock(index int64, missed bool) MissedBlock {
	return MissedBlock{
		Index:  index,
		Missed: missed,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params:       DefaultParams(),
		SigningInfos: []SigningInfos{},
		MissedBlocks: []ValidatorMissedBlocks{},
	}
}

// ValidateGenesis validates the slashing genesis parameters
func ValidateGenesis(data GenesisState) error {
	downtime := data.Params.SlashFractionDowntime
	if downtime.IsNegative() || downtime.GT(sdk.OneDec()) {
		return fmt.Errorf("slashing fraction downtime should be less than or equal to one and greater than zero, is %s", downtime.String())
	}

	dblSign := data.Params.SlashFractionDoubleSign
	if dblSign.IsNegative() || dblSign.GT(sdk.OneDec()) {
		return fmt.Errorf("slashing fraction double sign should be less than or equal to one and greater than zero, is %s", dblSign.String())
	}

	minSign := data.Params.MinSignedPerWindow
	if minSign.IsNegative() || minSign.GT(sdk.OneDec()) {
		return fmt.Errorf("min signed per window should be less than or equal to one and greater than zero, is %s", minSign.String())
	}

	downtimeJail := data.Params.DowntimeJailDuration
	if downtimeJail < 1*time.Minute {
		return fmt.Errorf("downtime unblond duration must be at least 1 minute, is %s", downtimeJail.String())
	}

	signedWindow := data.Params.SignedBlocksWindow
	if signedWindow < 10 {
		return fmt.Errorf("signed blocks window must be at least 10, is %d", signedWindow)
	}

	return nil
}
