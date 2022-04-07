package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host = "database"
	port = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) (Database, error) {
	db := Database{}

	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)

	conn, err := sql.Open("postgres", dns)

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
