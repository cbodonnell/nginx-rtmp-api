package main

import (
	"database/sql"
	"fmt"
	"log"
	
	// TODO: See about replacing with: https://github.com/jackc/pgx
	_ "github.com/lib/pq"
)

// db instance
var db *sql.DB

// connect to db
func connectDb(s DataSource) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.User, s.Password, s.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connecting to %s as %s\n", s.Dbname, s.User)
	return db
}

// ping db
func pingDb(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}