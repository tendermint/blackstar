package rest_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	rest2 "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/types/rest"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/testutil/network"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()
	cfg.NumValidators = 1

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestEncodeDecode() {
	val := s.network.Validators[0]

	// NOTE: this uses StdTx explicitly, don't migrate it!
	stdTx := authtypes.StdTx{
		Msgs: []sdk.Msg{&types.MsgSend{}},
		Fee: authtypes.StdFee{
			Amount: sdk.Coins{sdk.NewInt64Coin("foo", 10)},
			Gas:    10000,
		},
		Memo: "FOOBAR",
	}

	// NOTE: this uses amino explicitly, don't migrate it!
	cdc := val.ClientCtx.LegacyAmino

	bz, err := cdc.MarshalJSON(stdTx)
	s.Require().NoError(err)

	res, err := rest.PostRequest(fmt.Sprintf("%s/txs/encode", val.APIAddress), "application/json", bz)
	s.Require().NoError(err)

	var encodeResp rest2.EncodeResp
	err = cdc.UnmarshalJSON(res, &encodeResp)
	s.Require().NoError(err)

	bz, err = cdc.MarshalJSON(rest2.DecodeReq{Tx: encodeResp.Tx})
	s.Require().NoError(err)

	res, err = rest.PostRequest(fmt.Sprintf("%s/txs/decode", val.APIAddress), "application/json", bz)
	s.Require().NoError(err)

	var respWithHeight rest.ResponseWithHeight
	err = cdc.UnmarshalJSON(res, &respWithHeight)
	s.Require().NoError(err)
	var decodeResp rest2.DecodeResp
	err = cdc.UnmarshalJSON(respWithHeight.Result, &decodeResp)
	s.Require().NoError(err)
	s.Require().Equal(stdTx, authtypes.StdTx(decodeResp))
}

func (s *IntegrationTestSuite) TestBroadcastTxRequest() {
	// NOTE: this uses StdTx explicitly, don't migrate it!
	stdTx := authtypes.StdTx{
		Msgs: []sdk.Msg{&types.MsgSend{}},
		Fee: authtypes.StdFee{
			Amount: sdk.Coins{sdk.NewInt64Coin("foo", 10)},
			Gas:    10000,
		},
		Memo: "FOOBAR",
	}

	// we just test with async mode because this tx will fail - all we care about is that it got encoded and broadcast correctly
	res, err := s.broadcastReq(stdTx, "async")
	s.Require().NoError(err)
	var txRes sdk.TxResponse
	// NOTE: this uses amino explicitly, don't migrate it!
	s.Require().NoError(s.cfg.LegacyAmino.UnmarshalJSON(res, &txRes))
	// we just check for a non-empty TxHash here, the actual hash will depend on the underlying tx configuration
	s.Require().NotEmpty(txRes.TxHash)
}

func (s *IntegrationTestSuite) TestMultipleSignedBroadcastTxRequests() {

	// Set up TxConfig.
	aminoCdc := codec.NewLegacyAmino()
	// We're using TestMsg amino encoding in some tests, so register it here.
	txConfig := authtypes.StdTxConfig{Cdc: aminoCdc}
	txBuilder := txConfig.NewTxBuilder()

	val0 := s.network.Validators[0]
	val1 := s.network.Validators[0]
	msg := types.MsgSend{FromAddress: val0.Address, ToAddress: val1.Address, Amount: sdk.Coins{sdk.NewInt64Coin("foo", 100)}}

	feeAmount := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	gasLimit := testdata.NewTestGasLimit()
	txBuilder.SetMsgs(&msg)
	txBuilder.SetFeeAmount(feeAmount)
	txBuilder.SetGasLimit(gasLimit)

	txFactory := tx.Factory{}
	txFactory = txFactory.
		WithChainID(val0.ClientCtx.ChainID).
		WithKeybase(val0.ClientCtx.Keyring).
		WithTxConfig(val0.ClientCtx.TxConfig).
		WithSignMode(signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON)

	err := tx.Sign(txFactory, val0.Moniker, txBuilder)
	s.Require().NoError(err)

	stdTx := txBuilder.GetTx().(authtypes.StdTx)

	// we just test with async mode because this tx will fail - all we care about is that it got encoded and broadcast correctly
	res, err := s.broadcastReq(stdTx, "sync")
	s.Require().NoError(err)
	var txRes sdk.TxResponse
	// NOTE: this uses amino explicitly, don't migrate it!
	s.Require().NoError(s.cfg.LegacyAmino.UnmarshalJSON(res, &txRes))
	// we just check for a non-empty TxHash here, the actual hash will depend on the underlying tx configuration
	s.Require().Equal(txRes, sdk.TxResponse{})

}

func (s *IntegrationTestSuite) broadcastReq(stdTx authtypes.StdTx, mode string) ([]byte, error) {
	val := s.network.Validators[0]

	// NOTE: this uses amino explicitly, don't migrate it!
	cdc := val.ClientCtx.LegacyAmino

	req := rest2.BroadcastReq{
		Tx:   stdTx,
		Mode: mode,
	}
	bz, err := cdc.MarshalJSON(req)
	s.Require().NoError(err)

	return rest.PostRequest(fmt.Sprintf("%s/txs", val.APIAddress), "application/json", bz)
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
