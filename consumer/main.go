package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ornellast/bookstore/consumer/db"
)

var dbInstance db.Database

func main() {

	dbUser, dbPass, dbName := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")

	fmt.Printf("Params: %s -> %s -> %s\n", dbUser, dbPass, dbName)
	var dbErr error

	dbInstance, dbErr = db.Initialize(dbUser, dbPass, dbName)

	if dbErr != nil {
		log.Fatalf("Could not set up the database: %v", dbErr)
	}

	defer dbInstance.Conn.Close()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping Kafka Consumer")

}
