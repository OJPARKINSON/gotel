package run

import (
	"context"
	"log"
	"os"

	"github.com/OJPARKINSON/gotel/pkg/db"
	"github.com/OJPARKINSON/gotel/pkg/reservation"
	"github.com/OJPARKINSON/gotel/pkg/server"
	"github.com/OJPARKINSON/gotel/pkg/sqlite"
)

func Run(ctx context.Context) error {
	dbPath := os.Getenv("EASEL_DB_PATH")
	if dbPath == "" {
		dbPath = "easel.sqlite"
	}
	db, err := db.Open(dbPath, "easelSchema") // easelSchema)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rRepository := sqlite.NewReservationRepository(db)
	sRepository := reservation.NewService(rRepository)
	svr := server.NewServer(sRepository)

	svr.ListenAndServe(":8080")

	return nil
}
