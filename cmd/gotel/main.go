package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/OJPARKINSON/gotel/app/handlers"
	"github.com/OJPARKINSON/gotel/business/reservation"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		return err
	}
	defer pool.Close()

	resHandlers := &handlers.Reservation{Store: &reservation.Store{Pool: pool}}

	router := chi.NewRouter()
	router.Route("/reservations", func(r chi.Router) {
		r.Post("/", resHandlers.Create)
		r.Get("/{id}", resHandlers.GetByID)
	})

	server := &http.Server{
		Addr:         ":8011",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("listening on :8011")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(),
		10*time.Second)
	defer cancel()
	return server.Shutdown(shutdownCtx)

}
