package httpapi

import "net/http"

func NewRouter(db Pinger) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /health", healthHandler{db:db})
	return mux
}
