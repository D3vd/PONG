package middlewares

import (
	"net/http"
	"pong/log"
)

func RequestLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Request(log.RequestFields{
			Method:     r.Method,
			RequestURL: r.RequestURI,
			IP:         r.RemoteAddr,
		})
	}

	return http.HandlerFunc(fn)
}
