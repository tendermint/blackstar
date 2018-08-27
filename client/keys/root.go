package keys

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/gin-gonic/gin"
)

// Commands registers a sub-tree of commands to interact with
// local private key storage.
func Commands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Add or view local private keys",
		Long: `Keys allows you to manage your local keystore for tendermint.

    These keys may be in any format supported by go-crypto and can be
    used by light-clients, full nodes, or any other application that
    needs to sign with a private key.`,
	}
	cmd.AddCommand(
		addKeyCommand(),
		listKeysCmd,
		showKeysCmd,
		client.LineBreak,
		deleteKeyCommand(),
		updateKeyCommand(),
	)
	return cmd
}

// resgister REST routes
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/keys", QueryKeysRequestHandler).Methods("GET")
	r.HandleFunc("/keys", AddNewKeyRequestHandler).Methods("POST")
	r.HandleFunc("/keys/seed", SeedRequestHandler).Methods("GET")
	r.HandleFunc("/keys/{name}", GetKeyRequestHandler).Methods("GET")
	r.HandleFunc("/keys/{name}", UpdateKeyRequestHandler).Methods("PUT")
	r.HandleFunc("/keys/{name}", DeleteKeyRequestHandler).Methods("DELETE")
}

// RegisterSwaggerRoutes - Central function to define key management related routes that get registered by the main application
func RegisterSwaggerRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/keys", QueryKeysRequest)
	routerGroup.POST("/keys", AddNewKeyRequest)
	routerGroup.POST("/keys/:name/recover", RecoverResuest)
	routerGroup.GET("/keys/:name", GetKeyRequest)
	routerGroup.PUT("/keys/:name", UpdateKeyRequest)
	routerGroup.DELETE("/keys/:name", DeleteKeyRequest)
}