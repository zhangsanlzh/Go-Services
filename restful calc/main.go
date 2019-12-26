package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/{x}/add/{y}", addHandler).Methods("GET")
	r.HandleFunc("/api/{x}/sub/{y}", subHandler).Methods("GET")
	r.HandleFunc("/api/{x}/mult/{y}", multHandler).Methods("GET")
	r.HandleFunc("/api/{x}/div/{y}", divHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	type DATA struct {
		X int
		Y int
	}

	vars := mux.Vars(r)
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(vars["y"])

	rslt := DATA{
		X: x,
		Y: y,
	}
	if jsonStr, err := json.Marshal(rslt); err != nil {
		fmt.Println("error:", err)
	} else {
		os.Stdout.Write(jsonStr)
	}

	fmt.Fprintln(w, "x + y = ", x+y)
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	type DATA struct {
		X int
		Y int
	}

	vars := mux.Vars(r)
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(vars["y"])

	rslt := DATA{
		X: x,
		Y: y,
	}
	if jsonStr, err := json.Marshal(rslt); err != nil {
		fmt.Println("error:", err)
	} else {
		os.Stdout.Write(jsonStr)
	}

	fmt.Fprintln(w, "x - y = ", x-y)
}

func multHandler(w http.ResponseWriter, r *http.Request) {
	type DATA struct {
		X int
		Y int
	}

	vars := mux.Vars(r)
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(vars["y"])

	rslt := DATA{
		X: x,
		Y: y,
	}
	if jsonStr, err := json.Marshal(rslt); err != nil {
		fmt.Println("error:", err)
	} else {
		os.Stdout.Write(jsonStr)
	}

	fmt.Fprintln(w, "x * y = ", x*y)
}

func divHandler(w http.ResponseWriter, r *http.Request) {
	type DATA struct {
		X int
		Y int
	}

	vars := mux.Vars(r)
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(vars["y"])

	rslt := DATA{
		X: x,
		Y: y,
	}
	if jsonStr, err := json.Marshal(rslt); err != nil {
		fmt.Println("error:", err)
	} else {
		os.Stdout.Write(jsonStr)
	}

	fmt.Fprintln(w, "x / y = ", x/y)
}
