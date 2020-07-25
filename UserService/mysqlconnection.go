package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var sqldb *sql.DB

func GetSqlDB() *sql.DB {
	var err error
	dbUser := "root"
	dbPass := "Abhisql96()"
	dbName := "godb"
	sqldb, err := sql.Open("mysql", dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return sqldb
}
