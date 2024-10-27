package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/r4ndsen/sqlc-demo/internal/clipboard"
	"github.com/r4ndsen/sqlc-demo/internal/db"
	"github.com/r4ndsen/sqlc-demo/internal/server"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portEnv := os.Getenv("PORT")
	port := 3000
	if portEnv != "" {
		port, err = strconv.Atoi(portEnv)
		if err != nil {
			return fmt.Errorf("failed to parse PORT env variable: %w", err)
		}
	}

	ctx := context.Background()

	url, found := os.LookupEnv("DATABASE_URI_LOCAL")
	if !found {
		return fmt.Errorf("DATABASE_URI_LOCAL environment not set")
	}

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	go func() {
		s := server.New(queries, port)
		s.Start()
	}()

	cbm := clipboard.New(queries)
	cbm.Watch()

	select {}
}
