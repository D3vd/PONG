package routes

import (
	"pong/routes/ping"

	"github.com/go-chi/chi"
)

func Routes(r chi.Router) {
	r.Route("/ping", ping.Routes)
}
