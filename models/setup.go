package models

import (
	"database/sql"
	"fmt"
)

func ConnectDatabase(db_user string, db_pass string, db_name string) (*sql.DB, error) {
	fmt.Println("Connected To Database !")
	addr_sql := fmt.Sprintf("%s:%s@/%s", db_user, db_pass, db_name)
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", addr_sql)
	return db, err
}
