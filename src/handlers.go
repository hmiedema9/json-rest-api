/*
	Class to serve as the handlers for a JSON REST API configured in Go
	This uses the mux router from gorilla. In order to run, first install
	the router with 'go get -u github.com/gorilla/mux'
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func LanguageIndex(w http.ResponseWriter, r *http.Request) {
	languages := Languages{
		Language{Name: "Go"},
		Language{Name: "C#"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(languages); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(languages)
}

func LanguageShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	languageId := vars["languageId"]
	fmt.Fprintln(w, "Programming Language ID:", languageId)

}
