package sqlite

import (
	"context"
	"database/sql"

	"github.com/OJPARKINSON/gotel/pkg/reservation"
)

type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{
		db: db,
	}
}

func (r *ReservationRepository) List(ctx context.Context) ([]reservation.Reservation, error) {
	var reservations []reservation.Reservation

	err := r.db.QueryRow(`SELECT * FROM reservations`).Scan(&reservations)
	if err != nil {
		return nil, err
	}

	return reservations, nil
}
