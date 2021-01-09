package db

import (
	"database/sql"
	"fmt"
	"log"
	// import study
	_ "github.com/lib/pq"
)

const (
	// HOST is db host
	HOST = "database"
	// PORT is db port
	PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

// Database struct
type Database struct {
	Conn *sql.DB
}

// Initialize db connection
func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}
