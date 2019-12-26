package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/select/tbluser", tbluserHandler).Methods("GET")
	r.HandleFunc("/api/select/tbluser", tbluserHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
