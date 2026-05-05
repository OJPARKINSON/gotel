package resbus

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID         uuid.UUID
	HotelID    uuid.UUID
	RoomTypeID uuid.UUID
	StartDate  time.Time
	EndDate    time.Time
	Status     string
	GuestID    uuid.UUID
}

type NewReservation struct {
	HotelID    uuid.UUID
	RoomTypeID uuid.UUID
	StartDate  time.Time
	EndDate    time.Time
	GuestID    uuid.UUID
}
