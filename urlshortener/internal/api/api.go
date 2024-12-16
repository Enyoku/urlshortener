package api

import (
	"net/http"
	"urlShort/internal/storage"

	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
	db     *storage.DB
}

func New(db *storage.DB) *API {
	api := API{
		router: mux.NewRouter(),
		db:     db,
	}
	api.Endpoints()
	return &api
}

func (api *API) Endpoints() {
	api.router.Use(headerMiddleware)
}

func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}
