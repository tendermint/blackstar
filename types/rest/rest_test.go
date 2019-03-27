// Package rest provides HTTP types and primitives for REST
// requests validation and responses handling.
package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types"
)

type mockResponseWriter struct{}

func TestBaseReqValidateBasic(t *testing.T) {
	fromAddr := "cosmos1cq0sxam6x4l0sv9yz3a2vlqhdhvt2k6jtgcse0"
	tenstakes, err := types.ParseCoins("10stake")
	require.NoError(t, err)
	onestake, err := types.ParseDecCoins("1.0stake")
	require.NoError(t, err)

	req1 := NewBaseReq(
		fromAddr, "", "nonempty", "", "", 0, 0, tenstakes, nil, false,
	)
	req2 := NewBaseReq(
		"", "", "nonempty", "", "", 0, 0, tenstakes, nil, false,
	)
	req3 := NewBaseReq(
		fromAddr, "", "", "", "", 0, 0, tenstakes, nil, false,
	)
	req4 := NewBaseReq(
		fromAddr, "", "nonempty", "", "", 0, 0, tenstakes, onestake, false,
	)
	req5 := NewBaseReq(
		fromAddr, "", "nonempty", "", "", 0, 0, types.Coins{}, types.DecCoins{}, false,
	)

	tests := []struct {
		name string
		req  BaseReq
		w    http.ResponseWriter
		want bool
	}{
		{"ok", req1, httptest.NewRecorder(), true},
		{"neither fees nor gasprices provided", req5, httptest.NewRecorder(), true},
		{"empty from", req2, httptest.NewRecorder(), false},
		{"empty chain-id", req3, httptest.NewRecorder(), false},
		{"fees and gasprices provided", req4, httptest.NewRecorder(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.req.ValidateBasic(tt.w))
		})
	}
}

func TestParseHTTPArgs(t *testing.T) {
	req0, err := http.NewRequest("", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req1, err := http.NewRequest("", "/?limit=5", nil)
	if err != nil {
		t.Fatal(err)
	}
	req2, err := http.NewRequest("", "/?page=5", nil)
	if err != nil {
		t.Fatal(err)
	}
	req3, err := http.NewRequest("", "/?page=5&limit=5", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqE1, err := http.NewRequest("", "/?page=-1", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqE2, err := http.NewRequest("", "/?limit=-1", nil)
	if err != nil {
		t.Fatal(err)
	}

	req4, err := http.NewRequest("", "/?foo=faa", nil)
	if err != nil {
		t.Fatal(err)
	}
	err = req4.ParseForm()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name  string
		req   *http.Request
		w     http.ResponseWriter
		tags  []string
		page  int
		limit int
		err   bool
	}{
		{"no params", req0, httptest.NewRecorder(), []string{}, DefaultPage, DefaultLimit, false},
		{"Limit", req1, httptest.NewRecorder(), []string{}, DefaultPage, 5, false},
		{"Page", req2, httptest.NewRecorder(), []string{}, 5, DefaultLimit, false},
		{"Page and limit", req3, httptest.NewRecorder(), []string{}, 5, 5, false},

		{"error page 0", reqE1, httptest.NewRecorder(), []string{}, DefaultPage, DefaultLimit, true},
		{"error limit 0", reqE2, httptest.NewRecorder(), []string{}, DefaultPage, DefaultLimit, true},

		{"tags", req4, httptest.NewRecorder(), []string{"foo='faa'"}, DefaultPage, DefaultLimit, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tags, page, limit, err := ParseHTTPArgs(tt.req)
			if tt.err {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.tags, tags)
				require.Equal(t, tt.page, page)
				require.Equal(t, tt.limit, limit)
			}
		})
	}
}
