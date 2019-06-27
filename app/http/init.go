package http

import (
	"log"
	"net/http"

	"github.com/atedja/go-simple-http/app/config"
	"github.com/atedja/go-simple-http/app/http/routes"
	"github.com/atedja/go-simple-http/app/http/templates"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func init() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(config.Timeout))
	routes.Initialize(r)

	templates.Read("templates")

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t := templates.Load("helloworld.html")
		if t == nil {
			w.Header().Add("Content-Type", "text/plain")
			w.Write([]byte("Unable to load templates"))
			w.WriteHeader(500)
			return
		}

		w.Header().Add("Content-Type", "text/html")
		err := t.Execute(w, nil)
		if err != nil {
			w.WriteHeader(500)
		}
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Printf("Serving HTTP on port %s\n", config.Port)
	go http.ListenAndServe(config.Port, r)
}
