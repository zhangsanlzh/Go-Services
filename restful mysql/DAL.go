package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

func tbluserHandler(w http.ResponseWriter, r *http.Request) {
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
