package resapp

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/OJPARKINSON/gotel/business/domain/resbus"
	"github.com/google/uuid"
)

type Reservation struct {
	ID         string    `json:"id"`
	HotelID    string    `json:"hotelID"`
	RoomTypeID string    `json:"roomTypeID"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	Status     string    `json:"status"`
	GuestID    string    `json:"guestID"`
}

func (app Reservation) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

type NewReservation struct {
	HotelID    string    `json:"hotelID"`
	RoomTypeID string    `json:"roomTypeID"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	GuestID    string    `json:"guestID"`
}

func toAppReservation(resb resbus.Reservation) Reservation {
	return Reservation{
		ID:         resb.ID.String(),
		HotelID:    resb.HotelID.String(),
		RoomTypeID: resb.RoomTypeID.String(),
		StartDate:  resb.StartDate,
		EndDate:    resb.EndDate,
		GuestID:    resb.GuestID.String(),
	}
}

func toBusNewReservation(app NewReservation) (resbus.NewReservation, error) {
	hotelID, err := uuid.Parse(app.HotelID)
	if err != nil {
		return resbus.NewReservation{}, fmt.Errorf("parse hotelID: %w", err)
	}
	roomTypeID, err := uuid.Parse(app.RoomTypeID)
	if err != nil {
		return resbus.NewReservation{}, fmt.Errorf("parse roomTypeID: %w", err)
	}
	guestID, err := uuid.Parse(app.GuestID)
	if err != nil {
		return resbus.NewReservation{}, fmt.Errorf("parse guestID: %w", err)
	}

	return resbus.NewReservation{
		HotelID:    hotelID,
		RoomTypeID: roomTypeID,
		StartDate:  app.StartDate,
		EndDate:    app.EndDate,
		GuestID:    guestID,
	}, nil
}
