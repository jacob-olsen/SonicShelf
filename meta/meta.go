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
	db.Exec("CREATE TABLE IF NOT EXISTS voice(Id INTEGER, Name TEXT, addDate INTEGER, PRIMARY KEY(Id));")
	db.Exec("CREATE TABLE IF NOT EXISTS voiceSetings(Id INTEGER, Name TEXT, VoiceId INTEGER, PRIMARY KEY(Id), FOREIGN KEY(VoiceId) REFERENCES voice(Id));")

	//users
	db.Exec("CREATE TABLE IF NOT EXISTS user(Id INTEGER,UserName TEXT, voiceId INTEGER, PRIMARY KEY(Id), FOREIGN KEY(voiceId) REFERENCES voice(Id));")

	//auther
	db.Exec("CREATE TABLE IF NOT EXISTS auther(Id INTEGER,Name TEXT, PRIMARY KEY(Id));")

	//boogs
	db.Exec("CREATE TABLE IF NOT EXISTS book(Id INTEGER,Name TEXT, path TEXT, PRIMARY KEY(Id));")

	//pagers
	db.Exec("CREATE TABLE IF NOT EXISTS pageMeta(Id INTEGER,title TEXT, PRIMARY KEY(Id));")

	//key value pair
	db.Exec("CREATE TABLE IF NOT EXISTS keyValue(Id INTEGER, Key TEXT, Value TEXT, PRIMARY KEY(Id));")
	//db.Exec("CREATE TABLE IF NOT EXISTS page(Id INTEGER,title TEXT, PRIMARY KEY(Id));")
}

func SetKey(key string, value string) {
	if key == "" {
		fmt.Println("sql setKey:no key set")
		return
	}
	if value == "" {
		//remove item
		if GetKey(key) == "" {
			fmt.Println("sql setKey:remove non existen key")
		} else {
			db.Exec("DELETE FROM keyValue WHERE Key=?;", key)
		}
	} else {
		if GetKey(key) == "" {
			db.Exec("INSERT INTO keyValue (Key, Value) VALUES(?, ?);", key, value)
		} else {
			db.Exec("UPDATE keyValue SET Value=? WHERE Key=?;", value, key)
		}
	}

}
func GetKey(key string) string {
	value := ""
	db.QueryRow("SELECT Value FROM keyValue WHERE Key=?;", key).Scan(&value)
	return value
}

func AddVoice(name string) int {
	result, err := db.Exec("INSERT INTO voice (Name, addDate) VALUES(?, ?);", name, time.Now().Unix())
	if err != nil {
		print("sql addVoice:", err.Error())
	}
	id, _ := result.LastInsertId()
	return int(id)
}
func ListVoice() []Voice {
	result, err := db.Query("SELECT Id, Name, addDate FROM voice;")
	if err != nil {
		print("sql ListVoice:", err.Error())
	}
	var VoiceList []Voice
	for result.Next() {
		var newVoice Voice
		var uTime int64
		result.Scan(&newVoice.ID, &newVoice.Name, &uTime)
		newVoice.Addet = time.Unix(uTime, 0)
		VoiceList = append(VoiceList, newVoice)
	}
	return VoiceList
}
