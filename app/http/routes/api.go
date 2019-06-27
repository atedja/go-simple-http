package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func routeAPI() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"foo":"bar"}`))
	})
	return r
}
