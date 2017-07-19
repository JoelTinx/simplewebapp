package models

import (
	"log"
	"os"
	"encoding/json"
	"encoding/base64"
	"sync"
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type configuration struct {
	Host string
	Name string
	User string
	Pass string
	Port int
}

func init() {
	once.Do(func() {
		getCoonfigurationFile()
	})
	log.Println("SUCCESS: Config file loaded.")
}

var (
	Config *configuration
	once sync.Once
)

func getCoonfigurationFile() {
	file, err := os.Open("./app/config/db.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func getConnection() (*sql.DB, error) {
	getPassword := func() string {
		sDec, _ := base64.StdEncoding.DecodeString(Config.Pass)
		return string(sDec)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Config.User, getPassword(), Config.Host, Config.Port, Config.Name))
	if err != nil {
		log.Fatal(err.Error())
	}
	return db, err
}