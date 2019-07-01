package routes

import (
	"net/http"

	"github.com/atedja/go-simple-http/app/http/templates"
	"github.com/go-chi/chi"
)

func routeHello() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		err := templates.All.ExecuteTemplate(w, "hello", nil)
		if err != nil {
			w.Header().Add("Content-Type", "text/plain")
			w.Write([]byte("Unable to load templates"))
			w.WriteHeader(500)
			return
		}
	})
	return r
}
