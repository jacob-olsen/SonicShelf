package meta

import (
	"database/sql"
	"fmt"
	"os"
	"time"

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

	//piper voices
	db.Exec("CREATE TABLE IF NOT EXISTS voice(Id INTEGER,Name TEXT, addDate INTEGER, PRIMARY KEY(Id));")

	//users
	db.Exec("CREATE TABLE IF NOT EXISTS user(Id INTEGER,UserName TEXT, voiceId INTEGER, PRIMARY KEY(Id), FOREIGN KEY(voiceId) REFERENCES voice(Id));")

	//auther
	db.Exec("CREATE TABLE IF NOT EXISTS auther(Id INTEGER,Name TEXT, PRIMARY KEY(Id));")

	//boogs
	db.Exec("CREATE TABLE IF NOT EXISTS book(Id INTEGER,Name TEXT, path TEXT, PRIMARY KEY(Id));")

	//pagers
	db.Exec("CREATE TABLE IF NOT EXISTS pageMeta(Id INTEGER,title TEXT, PRIMARY KEY(Id));")
	//db.Exec("CREATE TABLE IF NOT EXISTS page(Id INTEGER,title TEXT, PRIMARY KEY(Id));")
}

func AddVoice(name string) int {
	result, err := db.Exec("INSERT INTO voice (Name, addDate) VALUES(?, ?);", name, time.Now().Unix())
	if err != nil {
		print("sql addVoice:", err.Error())
	}
	id, _ := result.LastInsertId()
	return int(id)
}
