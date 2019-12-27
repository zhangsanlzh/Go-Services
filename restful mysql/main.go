package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/select/tbluser", tbluser_All_Handler).Methods("GET")
	r.HandleFunc("/api/select/tbluser/{name}", tbluser_Single_Handler).Methods("GET")
	r.HandleFunc("/api/insert/tbluser/{name}/{age}", tbluser_Insert_Handler).Methods("POST")
	r.HandleFunc("/api/update/tbluser/{name}/{age}", tbluser_Update_Handler).Methods("PUT")
	r.HandleFunc("/api/delete/tbluser/{name}", tbluser_Delete_Handler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
