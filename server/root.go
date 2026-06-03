package server

import(
	"net/http"
	"encoding/json"
)


/*
func root(w, r) {
    set JSON header
    status 200
    encode a map with service info + routes
}
*/

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"service":"deadrop",
		"version":"1.0.0",
		"routes":[]string{
			"/health",
			"/",
		},
	})
}