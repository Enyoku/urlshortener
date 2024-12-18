package api

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	mathRand "math/rand/v2"
	"net/http"
	"urlShort/internal/models"
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

	api.router.HandleFunc("/", api.addUrl).Methods(http.MethodPost)
	api.router.HandleFunc("/{name}", api.redirect).Methods(http.MethodGet)
}

func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

// Adding url in db and response generated short url
func (api *API) addUrl(w http.ResponseWriter, r *http.Request) {
	var url models.Url

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shortUrl := "http://localhost:8080/" + randString(mathRand.IntN(8))

	urls := models.Urls{
		Url:      url.Url,
		ShortUrl: shortUrl,
	}

	storage.AddUrl(r.Context(), api.db, urls)

	resp, err := json.Marshal(models.Url{
		Url: shortUrl,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

// Redirect to original url in db by short url
func (api *API) redirect(w http.ResponseWriter, r *http.Request) {
	url := models.Url{Url: "http://" + r.Host + r.URL.Path}

	r_url, err := storage.GetUrl(r.Context(), api.db, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res string
	for _, u := range r_url {
		res = u.Url
	}

	http.Redirect(w, r, res, http.StatusPermanentRedirect)
}

// Generate random string for shorter url
func randString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	res := fmt.Sprintf("%x", b)
	return res
}
