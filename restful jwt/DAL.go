package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type TBLACCOUNT struct {
	Account  string `json:"account"`
	Password int    `json:"password,omitempty"`
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/testGo")
	if err != nil {
		log.Println(err.Error())
	}

	return db
}

func Valiad(account, password string) bool {
	con := Connect()
	defer con.Close()

	var valiad bool
	err := con.QueryRow("select if(count(*),'true','false') from tblAccount where account=? and password=?", account, password).Scan(&valiad)
	if err != nil {
		log.Println(err)
	}

	return valiad
}