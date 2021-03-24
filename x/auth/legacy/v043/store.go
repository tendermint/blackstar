package v043

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gogo/protobuf/grpc"
	"github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	delegatorDelegationPath = "/cosmos.staking.v1beta1.Query/DelegatorDelegations"
	balancesPath            = "/cosmos.bank.v1beta1.Query/AllBalances"
)

func migrateVestingAccounts(ctx sdk.Context, accounts []types.AccountI, queryServer grpc.Server) ([]types.AccountI, error) {
	for i := 0; i < len(accounts); i++ {
		asVesting, ok := accounts[i].(exported.VestingAccount)
		if !ok {
			continue
		}

		account := accounts[i]

		addr := account.GetAddress().String()
		balance, err := getBalance(
			ctx,
			addr,
			queryServer,
		)

		if err != nil {
			return nil, err
		}

		delegations, err := getDelegatorDelegationsSum(
			ctx,
			addr,
			queryServer,
		)

		if err != nil {
			return nil, err
		}

		asVesting, ok = resetVestingDelegatedBalances(asVesting)
		if !ok {
			continue
		}

		// balance before any delegation includes balance of delegation
		for _, coin := range delegations {
			balance = balance.Add(coin)
		}

		asVesting.TrackDelegation(ctx.BlockTime(), balance, delegations)

		accounts[i] = asVesting.(types.AccountI)
	}

	return accounts, nil
}

func resetVestingDelegatedBalances(evacct exported.VestingAccount) (exported.VestingAccount, bool) {
	// reset `DelegatedVesting` and `DelegatedFree` to zero
	df := sdk.NewCoins()
	dv := sdk.NewCoins()

	switch vacct := evacct.(type) {
	case *vestingtypes.ContinuousVestingAccount:
		vacct.DelegatedVesting = dv
		vacct.DelegatedFree = df
		return vacct, true
	case *vestingtypes.DelayedVestingAccount:
		vacct.DelegatedVesting = dv
		vacct.DelegatedFree = df
		return vacct, true
	case *vestingtypes.PeriodicVestingAccount:
		vacct.DelegatedVesting = dv
		vacct.DelegatedFree = df
		return vacct, true
	default:
		return nil, false
	}
}

func getDelegatorDelegationsSum(ctx sdk.Context, address string, queryServer grpc.Server) (sdk.Coins, error) {
	querier, ok := queryServer.(*baseapp.GRPCQueryRouter)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T wanted *baseapp.GRPCQueryRouter", queryServer)
	}

	queryFn := querier.Route(delegatorDelegationPath)

	q := &stakingtypes.QueryDelegatorDelegationsRequest{
		DelegatorAddr: address,
	}

	b, err := proto.Marshal(q)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal staking type query request, %w", err)
	}
	req := abci.RequestQuery{
		Data: b,
		Path: delegatorDelegationPath,
	}
	resp, err := queryFn(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("staking query error, %w", err)
	}

	balance := new(stakingtypes.QueryDelegatorDelegationsResponse)
	if err := proto.Unmarshal(resp.Value, balance); err != nil {
		return nil, fmt.Errorf("unable to unmarshal delegator query delegations: %w", err)
	}

	res := sdk.NewCoins()
	for _, i := range balance.DelegationResponses {
		res = res.Add(i.Balance)
	}

	return res, nil
}

func getBalance(ctx sdk.Context, address string, queryServer grpc.Server) (sdk.Coins, error) {
	querier, ok := queryServer.(*baseapp.GRPCQueryRouter)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T wanted *baseapp.GRPCQueryRouter", queryServer)
	}

	queryFn := querier.Route(balancesPath)

	q := &banktypes.QueryAllBalancesRequest{
		Address:    address,
		Pagination: nil,
	}
	b, err := proto.Marshal(q)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal bank type query request, %w", err)
	}

	req := abci.RequestQuery{
		Data: b,
		Path: balancesPath,
	}
	resp, err := queryFn(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("bank query error, %w", err)
	}
	balance := new(banktypes.QueryAllBalancesResponse)
	if err := proto.Unmarshal(resp.Value, balance); err != nil {
		return nil, fmt.Errorf("unable to unmarshal bank balance response: %w", err)
	}
	return balance.Balances, nil
}

// MigrateStore migrates vesting account to make the DelegatedVesting and DelegatedFree fields correctly
// track delegations.
// References: https://github.com/cosmos/cosmos-sdk/issues/8601, https://github.com/cosmos/cosmos-sdk/issues/8812
func MigrateStore(ctx sdk.Context, accounts []types.AccountI, queryServer grpc.Server) ([]types.AccountI, error) {
	return migrateVestingAccounts(ctx, accounts, queryServer)
}
