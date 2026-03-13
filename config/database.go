package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeDB() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	//MySQL user:passwrod@tcp(localhost:port)/databasename?parseTime=true/false
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Failed to open database (attempt: %d/5): %v", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			break
		}

		log.Printf("Failed to ping database (attempt: %d/5): %v", i+1, err)
		db.Close()
		continue
	}

	if err != nil {
		log.Fatal("Could not connect to database after 5 attempts: ", err)
	}

	_, err = db.Exec(fmt.Sprintf("USE %v", dbName))
	if err != nil {
		log.Fatalf("Unable to use database %v: %v", dbName, err)
	}

	log.Println("Database connected and Initialized Complete")
	return db
}
