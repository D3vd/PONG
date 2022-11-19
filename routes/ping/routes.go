package ping

import "github.com/go-chi/chi"

func Routes(r chi.Router) {
	r.Get("/", pingHandler)
}
