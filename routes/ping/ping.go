package ping

import (
	"encoding/json"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {

	resp := make(map[string]string)
	resp["msg"] = "pong"

	jsonResp, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(jsonResp)
}
