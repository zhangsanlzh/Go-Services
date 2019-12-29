package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func download_Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b, err := ioutil.ReadFile("images/" + vars["name"])
	if err != nil {
		w.Write(([]byte)("File Not Found"))
	}
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/download/{name}", download_Handler).Methods("GET")

	log.Println("serve is running")
	panic(http.ListenAndServe(":8000", r))
}
