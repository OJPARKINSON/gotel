package reservation

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Reservation struct {
	ID         uuid.UUID `json:"id"          db:"reservation_id"`
	HotelID    uuid.UUID `json:"hotelID"     db:"hotel_id"`
	RoomTypeID uuid.UUID `json:"roomTypeID"  db:"room_type_id"`
	StartDate  time.Time `json:"startDate"   db:"start_date"`
	EndDate    time.Time `json:"endDate"     db:"end_date"`
	Status     string    `json:"status"      db:"status"`
	GuestID    uuid.UUID `json:"guestID"     db:"guest_id"`
}

type NewReservation struct {
	HotelID    uuid.UUID `json:"hotelID"`
	RoomTypeID uuid.UUID `json:"roomTypeID"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	GuestID    uuid.UUID `json:"guestID"`
}

type Store struct{ Pool *pgxpool.Pool }

func (s *Store) Create(ctx context.Context, nr NewReservation) (Reservation, error) {
	r := Reservation{
		ID:         uuid.New(),
		HotelID:    nr.HotelID,
		RoomTypeID: nr.RoomTypeID,
		StartDate:  nr.StartDate,
		EndDate:    nr.EndDate,
		GuestID:    nr.GuestID,
		Status:     "pending",
	}

	const q = `
		INSERT INTO reservations
				(reservation_id, hotel_id, room_type_id, start_date, end_date,
				status, guest_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := s.Pool.Exec(ctx, q,
		r.ID, r.HotelID, r.RoomTypeID, r.StartDate, r.EndDate, r.Status,
		r.GuestID)
	if err != nil {
		return Reservation{}, fmt.Errorf("create reservation: %w", err)
	}
	return r, nil
}

func (s *Store) QueryByID(ctx context.Context, id uuid.UUID) (Reservation, error) {
	const q = `
		SELECT reservation_id, hotel_id, room_type_id, start_date,
			end_date, status, guest_id
		FROM reservations WHERE reservation_id = $1`

	rows, err := s.Pool.Query(ctx, q, id)
	if err != nil {
		return Reservation{}, nil
	}
	return pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[Reservation])
}
