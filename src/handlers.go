/*
	Class to serve as the handlers for a JSON REST API configured in Go
	This uses the mux router from gorilla. In order to run, first install
	the router with 'go get -u github.com/gorilla/mux'
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"io/ioutil"

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

// LanguageShow method for route /language/{languageId}
func LanguageShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	languageId := vars["languageId"]
	fmt.Fprintln(w, "Programming Language ID:", languageId)

}

// LanguageCreate method for POST /language
func LanguageCreate(w http.ResponseWriter, r *http.Request) {
	var language Language
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &language); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateLanguage(language)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
