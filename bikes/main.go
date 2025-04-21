package main

import (
	"context"
	"syscall"

	"log"
	"os"
	"os/signal"

	"time"

	"github.com/himanshuholmes/bikerental/bikes/server"
	"github.com/himanshuholmes/bikerental/db"
)

func main() {

	ctx, cancelFn := context.WithTimeout(context.Background(),time.Second*5)
	defer cancelFn()

	database,err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	srv,err := server.NewServer(ctx,database)
	if err !=nil {
		log.Fatalf("NewServer failed: %v", err)
	}
	srv.Run()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	signal := <-sigChan
	log.Printf("Received terminate, graceful shutdown %v",signal)
}