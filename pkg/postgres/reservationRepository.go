package postgres

import (
	"context"

	"github.com/OJPARKINSON/gotel/pkg/reservation"
	"github.com/jackc/pgx/v5"
)

type ReservationRepository struct {
	db *pgx.Conn
}

func NewReservationRepository(db *pgx.Conn) *ReservationRepository {
	return &ReservationRepository{
		db: db,
	}
}

func (r *ReservationRepository) List(ctx context.Context) ([]reservation.Reservation, error)
