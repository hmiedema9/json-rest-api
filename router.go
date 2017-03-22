/*
	Class to serve as the router for a JSON REST API configured in Go
	This uses the mux router from gorilla. In order to run, first install
	the router with 'go get -u github.com/gorilla/mux'
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Language struct {
	Name       string
	Experience int
	Difficulty int
}

type Languages []Language

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/Contact", Contact)
	router.HandleFunc("/Languages/{languageId}", LanguageIndex)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Contact Page!")
}

func LanguageIndex(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// languageId := vars["languageId"]
	// fmt.Fprintln(w, "Programming Language ID:", languageId)

	languages := Languages{
		Language{Name: "Go"},
		Language{Name: "C#"},
	}

	json.NewEncoder(w).Encode(languages)
}
