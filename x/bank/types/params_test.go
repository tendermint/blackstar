package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_validateSendEnabledParam(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"invalid type", args{sdk.NewCoin(sdk.DefaultBondDenom, sdk.OneInt())}, true},

		{"invalid empty denom send enabled", args{*NewSendEnabled("", true)}, true},
		{"invalid empty denom send disabled", args{*NewSendEnabled("", false)}, true},

		{"valid denom send enabled", args{*NewSendEnabled(sdk.DefaultBondDenom, true)}, false},
		{"valid denom send disabled", args{*NewSendEnabled(sdk.DefaultBondDenom, false)}, false},

		{"invalid denom send enabled", args{*NewSendEnabled("FOO", true)}, true},
		{"invalid denom send disabled", args{*NewSendEnabled("FOO", false)}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.wantErr, validateSendEnabled(tt.args.i) != nil)
		})
	}
}

func Test_sendParamEqual(t *testing.T) {
	paramsA := NewSendEnabled(sdk.DefaultBondDenom, true)
	paramsB := NewSendEnabled(sdk.DefaultBondDenom, true)
	paramsC := NewSendEnabled("foodenom", false)

	ok := paramsA.Equal(paramsB)
	require.True(t, ok)

	ok = paramsA.Equal(paramsC)
	require.False(t, ok)
}

func Test_sendParamString(t *testing.T) {
	paramString := "denom: foo\nenabled: false\n"
	param := NewSendEnabled("foo", false)

	require.Equal(t, paramString, param.String())
}

func Test_validateParams(t *testing.T) {
	params := DefaultParams()

	// default params have no error
	require.NoError(t, params.Validate())

	// default case is all denoms are enabled for sending
	require.True(t, params.IsSendEnabled(sdk.DefaultBondDenom))
	require.True(t, params.IsSendEnabled("foodenom"))

	params.DefaultSendEnabled = false
	params = params.SetSendEnabledParam("foodenom", true)

	require.NoError(t, validateSendEnabledParams(params.SendEnabled))
	require.True(t, params.IsSendEnabled("foodenom"))
	require.False(t, params.IsSendEnabled(sdk.DefaultBondDenom))

	params.DefaultSendEnabled = true
	params = params.SetSendEnabledParam("foodenom", false)

	require.NoError(t, validateSendEnabledParams(params.SendEnabled))
	require.False(t, params.IsSendEnabled("foodenom"))
	require.True(t, params.IsSendEnabled(sdk.DefaultBondDenom))

	params = params.SetSendEnabledParam("foodenom", true)
	require.True(t, params.IsSendEnabled("foodenom"))

	params = params.SetSendEnabledParam("foodenom", false)
	require.False(t, params.IsSendEnabled("foodenom"))

	require.True(t, params.IsSendEnabled("foodenom2"))
	params = params.SetSendEnabledParam("foodenom2", false)
	require.True(t, params.IsSendEnabled(""))
	require.True(t, params.IsSendEnabled(sdk.DefaultBondDenom))
	require.False(t, params.IsSendEnabled("foodenom2"))

	paramYaml := `send_enabled:
- denom: foodenom
  enabled: false
- denom: foodenom2
  enabled: false
defaultsendenabled: true
`
	require.Equal(t, paramYaml, params.String())

	params = NewParams(true, SendEnabledParams{
		NewSendEnabled("foodenom", false),
		NewSendEnabled("foodenom", true), // this is not allowed
	})

	// fails due to duplicate entries.
	require.Error(t, params.Validate())

	// fails due to invalid type
	require.Error(t, validateSendEnabledParams(NewSendEnabled("foodenom", true)))

	require.Error(t, validateSendEnabledParams(SendEnabledParams{NewSendEnabled("INVALIDDENOM", true)}))
}
