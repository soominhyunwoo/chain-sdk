package rest

import (
	"github.com/gorilla/mux"

	"github.com/soominhyunwoo/chain-sdk/client"
	"github.com/soominhyunwoo/chain-sdk/client/rest"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
}
