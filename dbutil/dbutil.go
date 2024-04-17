package dbutil

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"

)

func OpenDB () {
	dbFile := "rbac.db"
	var err error
	DB, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err.Error())
	}
}

var DB *sql.DB