package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {

	var err error
	dsn := LoadConfig().DB_URL
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("error occured while connecting to the db")
		return
	}

	if err := DB.Ping(); err != nil {
		log.Println("error occured while pinigng to the db")
		return
	}

	fmt.Println("connected to the db")

}
