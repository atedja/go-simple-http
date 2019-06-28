package http

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "public")
	FileServer(r, "/", http.Dir(filesDir))

	log.Printf("Serving HTTP on port %s\n", config.Port)
	go http.ListenAndServe(config.Port, r)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
