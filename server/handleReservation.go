package server

import (
	"encoding/json"
	"net/http"

	"github.com/OJPARKINSON/gotel/reservation"
)

type ReservationHandler struct {
	service *reservation.Service
}

func (h *ReservationHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
