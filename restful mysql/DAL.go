package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type TBLUSER struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/testGo")
	if err != nil {
		log.Println(err.Error())
	}

	return db
}

func tbluser_All_Handler(w http.ResponseWriter, r *http.Request) {
	con := Connect()
	defer con.Close()

	rows, err := con.Query("select * from tbluser")
	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	for rows.Next() {
		var user TBLUSER
		rows.Scan(&user.Name, &user.Age)
		content, _ := json.Marshal(user)
		fmt.Fprintln(w, string(content))
	}
}

func tbluser_Single_Handler(w http.ResponseWriter, r *http.Request) {
	con := Connect()
	defer con.Close()

	vars := mux.Vars(r)
	rows, err := con.Query("select * from tbluser where name = ?", vars["name"])
	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	for rows.Next() {
		var user TBLUSER
		rows.Scan(&user.Name, &user.Age)
		content, _ := json.Marshal(user)
		fmt.Fprintln(w, string(content))
	}
}

func tbluser_Insert_Handler(w http.ResponseWriter, r *http.Request) {
	con := Connect()
	defer con.Close()

	vars := mux.Vars(r)
	con.Exec("insert into tbluser values(?, ?)", vars["name"], vars["age"])

	w.WriteHeader(http.StatusOK)
}

func tbluser_Update_Handler(w http.ResponseWriter, r *http.Request) {
	con := Connect()
	defer con.Close()

	vars := mux.Vars(r)
	con.Exec("update tbluser set age = ? where name = ?", vars["age"], vars["name"])

	w.WriteHeader(http.StatusOK)
}

func tbluser_Delete_Handler(w http.ResponseWriter, r *http.Request) {
	con := Connect()
	defer con.Close()

	vars := mux.Vars(r)
	con.Exec("delete from tbluser where name =?", vars["name"])

	w.WriteHeader(http.StatusOK)
}
