// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	"github.com/cosmos/cosmos-sdk/server/rosetta/client/sdk/types"
	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/cosmos/cosmos-sdk/client/rpc"
)

// SdkClient is an autogenerated mock type for the SdkClient type
type SdkClient struct {
	mock.Mock
}

// GetAuthAccount provides a mock function with given fields: ctx, address, height
func (_m *SdkClient) GetAuthAccount(ctx context.Context, address string, height int64) (types.AccountResponse, error) {
	ret := _m.Called(ctx, address, height)

	var r0 types.AccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) types.AccountResponse); ok {
		r0 = rf(ctx, address, height)
	} else {
		r0 = ret.Get(0).(types.AccountResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, address, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeInfo provides a mock function with given fields: ctx
func (_m *SdkClient) GetNodeInfo(ctx context.Context) (rpc.NodeInfoResponse, error) {
	ret := _m.Called(ctx)

	var r0 rpc.NodeInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context) rpc.NodeInfoResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(rpc.NodeInfoResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTx provides a mock function with given fields: ctx, hash
func (_m *SdkClient) GetTx(ctx context.Context, hash string) (cosmos_sdktypes.TxResponse, error) {
	ret := _m.Called(ctx, hash)

	var r0 cosmos_sdktypes.TxResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) cosmos_sdktypes.TxResponse); ok {
		r0 = rf(ctx, hash)
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.TxResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostTx provides a mock function with given fields: ctx, bytes
func (_m *SdkClient) PostTx(ctx context.Context, bytes []byte) (cosmos_sdktypes.TxResponse, error) {
	ret := _m.Called(ctx, bytes)

	var r0 cosmos_sdktypes.TxResponse
	if rf, ok := ret.Get(0).(func(context.Context, []byte) cosmos_sdktypes.TxResponse); ok {
		r0 = rf(ctx, bytes)
	} else {
		r0 = ret.Get(0).(cosmos_sdktypes.TxResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, bytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
