package routes

import (
	"github.com/go-chi/chi"
)

func Initialize(r chi.Router) {
	r.Mount("/api", routeAPI())
}
