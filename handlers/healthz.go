package gabookApiHandlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RouteSetupHealthz sets up Healthz routes
func RouteSetupHealthz(s *mux.Router) {
	s.HandleFunc("/healthz", handlerHealthz).Methods("GET")
}

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{ "status": "ok" }`)
}
