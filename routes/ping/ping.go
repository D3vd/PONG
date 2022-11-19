package ping

import (
	"net/http"
	"pong/log"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("ping handler")
}
