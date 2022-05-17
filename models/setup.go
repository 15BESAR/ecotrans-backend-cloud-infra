package models

import (
	"database/sql"
	"fmt"
)

func ConnectDatabase(dbUser string, dbPwd string, dbName string, dbTCPHost string, dbPort string) (*sql.DB, error) {
	addr_sql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPwd, dbTCPHost, dbPort, dbName)
	fmt.Println(dbPort)
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", addr_sql)
	return db, err
}
