package server

import (
	"net/http"

	"github.com/OJPARKINSON/gotel/pkg/reservation"
)

type Server struct {
	reservationHandler *ReservationHandler
}

func NewServer(reservationService *reservation.Service) *Server {
	return &Server{
		reservationHandler: &ReservationHandler{service: reservationService},
	}
}

func (s *Server) ListenAndServe(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /reservations", s.reservationHandler.List)
	return http.ListenAndServe(addr, mux)
}
