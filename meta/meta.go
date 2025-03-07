package meta

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Setup() {
	var err error
	db, err = sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("sql is online")
}
