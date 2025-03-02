package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers ibcdex-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/ibcdex/slogs/{id}", getSlogHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/ibcdex/slogs", listSlogHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/ibcdex/slogs", createSlogHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibcdex/slogs/{id}", updateSlogHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/ibcdex/slogs/{id}", deleteSlogHandler(clientCtx)).Methods("POST")

}
