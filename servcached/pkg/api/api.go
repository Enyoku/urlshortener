package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

type API struct {
	router *mux.Router
	db     *redis.Client
}

func New(client *redis.Client) *API {
	api := API{
		router: mux.NewRouter(),
		db:     client,
	}
	api.Endpoints()
	return &api
}

func (api *API) Run(port int) {
	p := ":" + strconv.Itoa(port)
	http.ListenAndServe(p, api.router)
}

func (api *API) Endpoints() {
	api.router.HandleFunc("/", api.hiHandler).Methods(http.MethodGet)
}

func (api *API) hiHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]int, 2)
	data["status"] = http.StatusOK

	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}
