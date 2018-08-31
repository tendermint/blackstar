package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/gin-gonic/gin"
)

// registerSwaggerQueryRoutes - Central function to define account query related routes that get registered by the main application
func registerSwaggerQueryRoutes(routerGroup *gin.RouterGroup, ctx context.CLIContext, cdc *wire.Codec, storeName string) {
	routerGroup.GET("/bank/balances/:account", queryAccountRequestHandler(storeName, cdc, authcmd.GetAccountDecoder(cdc), ctx))
}

func queryAccountRequestHandler(storeName string, cdc *wire.Codec, decoder auth.AccountDecoder, ctx context.CLIContext) gin.HandlerFunc {
	return func(gtx *gin.Context) {

		bech32addr := gtx.Param("account")

		addr, err := sdk.AccAddressFromBech32(bech32addr)
		if err != nil {
			utils.NewError(gtx, http.StatusConflict, err)
			return
		}

		res, err := ctx.QueryStore(auth.AddressStoreKey(addr), storeName)
		if err != nil {
			utils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't query account. Error: %s", err.Error()))
			return
		}

		// the query will return empty if there is no data for this account
		if len(res) == 0 {
			utils.NewError(gtx, http.StatusNoContent, fmt.Errorf("this account info is nil+"))
			return
		}

		// decode the value
		account, err := decoder(res)
		if err != nil {
			utils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't parse query result. Result: %s. Error: %s", res, err.Error()))
			return
		}

		// print out whole account
		output, err := cdc.MarshalJSON(account.GetCoins())
		if err != nil {
			utils.NewError(gtx, http.StatusInternalServerError, fmt.Errorf("couldn't marshall query result. Error: %s", err.Error()))
			return
		}

		utils.NormalResponse(gtx, output)
	}
}
