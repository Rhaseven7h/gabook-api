package main

import (
	"log"
	"net/http"

	"github.com/Rhaseven7h/gabook-api/database"
	"github.com/Rhaseven7h/gabook-api/handlers"
	"github.com/gorilla/mux"
)

const (
	listenAddress = "0.0.0.0:6500"
	logsPrefix    = "GABOOKS\u0020"
)

func routePrepare(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func setupRoutes() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/v1/").Subrouter()
	gabookApiHandlers.RouteSetupHealthz(s)
	gabookApiHandlers.RouteSetupAuthors(s)
	return r
}

func logRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	log.SetPrefix(logsPrefix)
	_ = gabookApiDatabase.GetGaBookDB()
	//defer gbdb.Close()
	log.Println("GaBooks eBook Management System API")
	log.Println("Listening on " + listenAddress)
	routes := setupRoutes()

	http.ListenAndServe(listenAddress, logRequests(routes))
}
