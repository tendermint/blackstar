package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// module name
	ModuleName = "nft"

	// StoreKey is the default store key for mint
	StoreKey = ModuleName
)

var (
	NFTKey        = []byte{0x01} // key for a nft
	NFTByOwnerKey = []byte{0x02} // prefix for each key for an nft id
)

// GetNFTKey creates the key for nft
// VALUE: nft
func GetNFTKey(id string) []byte {
	return append(NFTKey, GetNFTIDBytes(id)...)
}

// GetNFTByOwnerKey creates the key for nft id
// VALUE: nft id
func GetNFTByOwnerKey(owner sdk.AccAddress) []byte {
	return append(NFTByOwnerKey, address.MustLengthPrefix(owner)...)
}

// GetNFTIDBytes returns the byte representation of the nftID
func GetNFTIDBytes(id string) []byte {
	// TODO
	return []byte(id)
}

// GetNFTID returns nftID by the byte representation of the nftID
func GetNFTID(idBytes []byte) string {
	// TODO
	return string(idBytes)
}
