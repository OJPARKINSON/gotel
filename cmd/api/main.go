package main

import (
	"context"
	"fmt"
	"os"

	"github.com/OJPARKINSON/gotel/postgres"
	"github.com/OJPARKINSON/gotel/reservation"
	"github.com/OJPARKINSON/gotel/server"
	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rRepository := postgres.NewReservationRepository(conn)
	sRepository := reservation.NewService(rRepository)
	svr := server.NewServer(sRepository)

	svr.ListenAndServe(":8080")
}
