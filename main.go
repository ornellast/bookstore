package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ornellast/bucketeer/db"
	"github.com/ornellast/bucketeer/handlers"
)

func main() {
	const addr = ":8080"
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	dbUser, dbPass, dbName := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")

	database, err := db.Initialize(dbUser, dbPass, dbName)

	if err != nil {
		log.Fatalf("Could not set up the database: %v", err)
	}

	defer database.Conn.Close()

	httHandler := handlers.NewHandler(database)

	server := &http.Server{
		Handler: httHandler,
	}

	go func() {
		server.Serve(listener)
	}()

	defer Stop(server)

	log.Printf("Started server on %s", addr)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server")

}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down the server correctly: %v", err)
		os.Exit(1)
	}
}
