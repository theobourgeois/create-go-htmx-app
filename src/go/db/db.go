package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)


var db * sql.DB

func InitDB() {
	cfg := mysql.Config{
		User:  	"root",
		Passwd: "root", 
		Net:    "tcp",
		Addr:   "127.0.0.1:8889", 
		DBName: "trackrack",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
			log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
			log.Fatal(pingErr)
	}
	fmt.Println("Database Connected Successfully on port 8889...")
}

func Query(query string, args ...any) (*sql.Rows, error) {
	return db.Query(query)
}