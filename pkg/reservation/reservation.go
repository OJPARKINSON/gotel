package reservation

import (
	"time"

	"github.com/google/uuid"
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
