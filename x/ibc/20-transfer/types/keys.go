package types

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctypes "github.com/cosmos/cosmos-sdk/x/ibc/types"
)

const (
	// SubModuleName defines the IBC transfer name
	SubModuleName = "transfer"

	// StoreKey is the store key string for IBC transfer
	StoreKey = SubModuleName

	// RouterKey is the message route for IBC transfer
	RouterKey = SubModuleName

	// QuerierRoute is the querier route for IBC transfer
	QuerierRoute = SubModuleName
)

// ModuleAccountName defines the IBC transfer module account name
var ModuleAccountName = fmt.Sprintf("%s/%s", ibctypes.ModuleName, SubModuleName)

// GetEscrowAddress returns the escrow address for the specified channel
func GetEscrowAddress(chanID string) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(chanID)))
}

// GetDenomPrefix returns the receiving denomination prefix
func GetDenomPrefix(portID, channelID string) string {
	return fmt.Sprintf("%s/%s", portID, channelID)
}
