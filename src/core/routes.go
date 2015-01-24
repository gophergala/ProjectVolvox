package core

import (
	"github.com/gorilla/mux"
	"net/http"
)

func init() {

	rmux := mux.NewRouter()

	// I expect someone else on the team to know what routing scheme we want to use.
	http.Handle("/jackman/", rmux)

	rmux.HandleFunc("/submit/", submitHandler)
}
