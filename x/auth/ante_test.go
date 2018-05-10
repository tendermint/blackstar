package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/abci/types"
	crypto "github.com/tendermint/go-crypto"
	"github.com/tendermint/tmlibs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
)

func newTestMsg(addrs ...bam.Address) *bam.TestMsg {
	return bam.NewTestMsg(addrs...)
}

func newStdFee() bam.StdFee {
	return bam.NewStdFee(100,
		bam.Coin{"atom", 150},
	)
}

// coins to more than cover the fee
func newCoins() bam.Coins {
	return bam.Coins{
		{"atom", 10000000},
	}
}

// generate a priv key and return it with its address
func privAndAddr() (crypto.PrivKey, bam.Address) {
	priv := crypto.GenPrivKeyEd25519()
	addr := priv.PubKey().Address()
	return priv, addr
}

// run the tx through the anteHandler and ensure its valid
func checkValidTx(t *testing.T, anteHandler bam.AnteHandler, ctx sdk.Context, tx bam.Tx) {
	_, result, abort := anteHandler(ctx, tx)
	assert.False(t, abort)
	assert.Equal(t, bam.ABCICodeOK, result.Code)
	assert.True(t, result.IsOK())
}

// run the tx through the anteHandler and ensure it fails with the given code
func checkInvalidTx(t *testing.T, anteHandler bam.AnteHandler, ctx sdk.Context, tx bam.Tx, code bam.CodeType) {
	_, result, abort := anteHandler(ctx, tx)
	assert.True(t, abort)
	assert.Equal(t, bam.ToABCICode(bam.CodespaceRoot, code), result.Code)
}

func newTestTx(ctx sdk.Context, msg bam.Msg, privs []crypto.PrivKey, seqs []int64, fee bam.StdFee) bam.Tx {
	signBytes := bam.StdSignBytes(ctx.ChainID(), seqs, fee, msg)
	return newTestTxWithSignBytes(msg, privs, seqs, fee, signBytes)
}

func newTestTxWithSignBytes(msg bam.Msg, privs []crypto.PrivKey, seqs []int64, fee bam.StdFee, signBytes []byte) bam.Tx {
	sigs := make([]bam.StdSignature, len(privs))
	for i, priv := range privs {
		sigs[i] = bam.StdSignature{PubKey: priv.PubKey(), Signature: priv.Sign(signBytes), Sequence: seqs[i]}
	}
	tx := bam.NewStdTx(msg, fee, sigs)
	return tx
}

// Test various error cases in the AnteHandler control flow.
func TestAnteHandlerSigErrors(t *testing.T) {
	// setup
	ms, capKey := setupMultiStore()
	cdc := wire.NewCodec()
	RegisterBaseAccount(cdc)
	mapper := NewAccountMapper(cdc, capKey, &BaseAccount{})
	anteHandler := NewAnteHandler(mapper, BurnFeeHandler)
	ctx := bam.NewContext(ms, abci.Header{ChainID: "mychainid"}, false, nil, log.NewNopLogger())

	// keys and addresses
	priv1, addr1 := privAndAddr()
	priv2, addr2 := privAndAddr()

	// msg and signatures
	var tx bam.Tx
	msg := newTestMsg(addr1, addr2)
	fee := newStdFee()

	// test no signatures
	privs, seqs := []crypto.PrivKey{}, []int64{}
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeUnauthorized)

	// test num sigs dont match GetSigners
	privs, seqs = []crypto.PrivKey{priv1}, []int64{0}
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeUnauthorized)

	// test an unrecognized account
	privs, seqs = []crypto.PrivKey{priv1, priv2}, []int64{0, 0}
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeUnknownAddress)

	// save the first account, but second is still unrecognized
	acc1 := mapper.NewAccountWithAddress(ctx, addr1)
	acc1.SetCoins(fee.Amount)
	mapper.SetAccount(ctx, acc1)
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeUnknownAddress)
}

// Test logic around sequence checking with one signer and many signers.
func TestAnteHandlerSequences(t *testing.T) {
	// setup
	ms, capKey := setupMultiStore()
	cdc := wire.NewCodec()
	RegisterBaseAccount(cdc)
	mapper := NewAccountMapper(cdc, capKey, &BaseAccount{})
	anteHandler := NewAnteHandler(mapper, BurnFeeHandler)
	ctx := bam.NewContext(ms, abci.Header{ChainID: "mychainid"}, false, nil, log.NewNopLogger())

	// keys and addresses
	priv1, addr1 := privAndAddr()
	priv2, addr2 := privAndAddr()

	// set the accounts
	acc1 := mapper.NewAccountWithAddress(ctx, addr1)
	acc1.SetCoins(newCoins())
	mapper.SetAccount(ctx, acc1)
	acc2 := mapper.NewAccountWithAddress(ctx, addr2)
	acc2.SetCoins(newCoins())
	mapper.SetAccount(ctx, acc2)

	// msg and signatures
	var tx bam.Tx
	msg := newTestMsg(addr1)
	fee := newStdFee()

	// test good tx from one signer
	privs, seqs := []crypto.PrivKey{priv1}, []int64{0}
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkValidTx(t, anteHandler, ctx, tx)

	// test sending it again fails (replay protection)
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeInvalidSequence)

	// fix sequence, should pass
	seqs = []int64{1}
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkValidTx(t, anteHandler, ctx, tx)

	// new tx with another signer and correct sequences
	msg = newTestMsg(addr1, addr2)
	privs, seqs = []crypto.PrivKey{priv1, priv2}, []int64{2, 0}
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkValidTx(t, anteHandler, ctx, tx)

	// replay fails
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeInvalidSequence)

	// tx from just second signer with incorrect sequence fails
	msg = newTestMsg(addr2)
	privs, seqs = []crypto.PrivKey{priv2}, []int64{0}
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeInvalidSequence)

	// fix the sequence and it passes
	tx = newTestTx(ctx, msg, []crypto.PrivKey{priv2}, []int64{1}, fee)
	checkValidTx(t, anteHandler, ctx, tx)

	// another tx from both of them that passes
	msg = newTestMsg(addr1, addr2)
	privs, seqs = []crypto.PrivKey{priv1, priv2}, []int64{3, 2}
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkValidTx(t, anteHandler, ctx, tx)
}

// Test logic around fee deduction.
func TestAnteHandlerFees(t *testing.T) {
	// setup
	ms, capKey := setupMultiStore()
	cdc := wire.NewCodec()
	RegisterBaseAccount(cdc)
	mapper := NewAccountMapper(cdc, capKey, &BaseAccount{})
	anteHandler := NewAnteHandler(mapper, BurnFeeHandler)
	ctx := bam.NewContext(ms, abci.Header{ChainID: "mychainid"}, false, nil, log.NewNopLogger())

	// keys and addresses
	priv1, addr1 := privAndAddr()

	// set the accounts
	acc1 := mapper.NewAccountWithAddress(ctx, addr1)
	mapper.SetAccount(ctx, acc1)

	// msg and signatures
	var tx bam.Tx
	msg := newTestMsg(addr1)
	privs, seqs := []crypto.PrivKey{priv1}, []int64{0}
	fee := bam.NewStdFee(100,
		bam.Coin{"atom", 150},
	)

	// signer does not have enough funds to pay the fee
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeInsufficientFunds)

	acc1.SetCoins(bam.Coins{{"atom", 149}})
	mapper.SetAccount(ctx, acc1)
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeInsufficientFunds)

	acc1.SetCoins(bam.Coins{{"atom", 150}})
	mapper.SetAccount(ctx, acc1)
	checkValidTx(t, anteHandler, ctx, tx)
}

func TestAnteHandlerBadSignBytes(t *testing.T) {
	// setup
	ms, capKey := setupMultiStore()
	cdc := wire.NewCodec()
	RegisterBaseAccount(cdc)
	mapper := NewAccountMapper(cdc, capKey, &BaseAccount{})
	anteHandler := NewAnteHandler(mapper, BurnFeeHandler)
	ctx := bam.NewContext(ms, abci.Header{ChainID: "mychainid"}, false, nil, log.NewNopLogger())

	// keys and addresses
	priv1, addr1 := privAndAddr()
	priv2, addr2 := privAndAddr()

	// set the accounts
	acc1 := mapper.NewAccountWithAddress(ctx, addr1)
	acc1.SetCoins(newCoins())
	mapper.SetAccount(ctx, acc1)
	acc2 := mapper.NewAccountWithAddress(ctx, addr2)
	acc2.SetCoins(newCoins())
	mapper.SetAccount(ctx, acc2)

	var tx bam.Tx
	msg := newTestMsg(addr1)
	fee := newStdFee()
	fee2 := newStdFee()
	fee2.Gas += 100
	fee3 := newStdFee()
	fee3.Amount[0].Amount += 100

	// test good tx and signBytes
	privs, seqs := []crypto.PrivKey{priv1}, []int64{0}
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkValidTx(t, anteHandler, ctx, tx)

	chainID := ctx.ChainID()
	chainID2 := chainID + "somemorestuff"
	codeUnauth := bam.CodeUnauthorized

	cases := []struct {
		chainID string
		seqs    []int64
		fee     bam.StdFee
		msg     bam.Msg
		code    bam.CodeType
	}{
		{chainID2, []int64{1}, fee, msg, codeUnauth},               // test wrong chain_id
		{chainID, []int64{2}, fee, msg, codeUnauth},                // test wrong seqs
		{chainID, []int64{1, 2}, fee, msg, codeUnauth},             // test wrong seqs
		{chainID, []int64{1}, fee, newTestMsg(addr2), codeUnauth},  // test wrong msg
		{chainID, []int64{1}, fee2, newTestMsg(addr2), codeUnauth}, // test wrong fee
		{chainID, []int64{1}, fee3, newTestMsg(addr2), codeUnauth}, // test wrong fee
	}

	privs, seqs = []crypto.PrivKey{priv1}, []int64{1}
	for _, cs := range cases {
		tx := newTestTxWithSignBytes(
			msg, privs, seqs, fee,
			bam.StdSignBytes(cs.chainID, cs.seqs, cs.fee, cs.msg),
		)
		checkInvalidTx(t, anteHandler, ctx, tx, cs.code)
	}

	// test wrong signer if public key exist
	privs, seqs = []crypto.PrivKey{priv2}, []int64{1}
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeUnauthorized)

	// test wrong signer if public doesn't exist
	msg = newTestMsg(addr2)
	privs, seqs = []crypto.PrivKey{priv1}, []int64{0}
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeInvalidPubKey)

}

func TestAnteHandlerSetPubKey(t *testing.T) {
	// setup
	ms, capKey := setupMultiStore()
	cdc := wire.NewCodec()
	RegisterBaseAccount(cdc)
	mapper := NewAccountMapper(cdc, capKey, &BaseAccount{})
	anteHandler := NewAnteHandler(mapper, BurnFeeHandler)
	ctx := bam.NewContext(ms, abci.Header{ChainID: "mychainid"}, false, nil, log.NewNopLogger())

	// keys and addresses
	priv1, addr1 := privAndAddr()
	_, addr2 := privAndAddr()

	// set the accounts
	acc1 := mapper.NewAccountWithAddress(ctx, addr1)
	acc1.SetCoins(newCoins())
	mapper.SetAccount(ctx, acc1)
	acc2 := mapper.NewAccountWithAddress(ctx, addr2)
	acc2.SetCoins(newCoins())
	mapper.SetAccount(ctx, acc2)

	var tx bam.Tx

	// test good tx and set public key
	msg := newTestMsg(addr1)
	privs, seqs := []crypto.PrivKey{priv1}, []int64{0}
	fee := newStdFee()
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkValidTx(t, anteHandler, ctx, tx)

	acc1 = mapper.GetAccount(ctx, addr1)
	require.Equal(t, acc1.GetPubKey(), priv1.PubKey())

	// test public key not found
	msg = newTestMsg(addr2)
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	sigs := tx.GetSignatures()
	sigs[0].PubKey = nil
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeInvalidPubKey)

	acc2 = mapper.GetAccount(ctx, addr2)
	assert.Nil(t, acc2.GetPubKey())

	// test invalid signature and public key
	tx = newTestTx(ctx, msg, privs, seqs, fee)
	checkInvalidTx(t, anteHandler, ctx, tx, bam.CodeInvalidPubKey)

	acc2 = mapper.GetAccount(ctx, addr2)
	assert.Nil(t, acc2.GetPubKey())
}
