package gabookApiHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rhaseven7h/gabook-api/database"
	"github.com/Rhaseven7h/gabook-api/models"
	"github.com/gorilla/mux"
)

// RouteSetupAuthors sets up Healthz routes
func RouteSetupAuthors(s *mux.Router) {
	s.HandleFunc("/authors", handlerAuthorsList).Methods("GET")
}

func handlerAuthorsList(w http.ResponseWriter, r *http.Request) {
	authorsCol := gabookApiDatabase.GetGaBookDB().Database.C("authors")

	authors := make([]gabookApiModels.Author, 100)
	err := authorsCol.Find(nil).Limit(100).All(&authors)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"error": "database error - %v"}`, err)
	} else {
		if bson, err := json.Marshal(map[string][]gabookApiModels.Author{"authors": authors}); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, `{"error": "unable to encode - %s"}`, err)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bson)
		}
	}
}
