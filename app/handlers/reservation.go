package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/OJPARKINSON/gotel/business/reservation"
	"github.com/google/uuid"
)

type Reservation struct {
	Store *reservation.Store
}

func (h *Reservation) Create(w http.ResponseWriter, r *http.Request) {
	var nr reservation.NewReservation
	if err := json.NewDecoder(r.Body).Decode(&nr); err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.Store.Create(r.Context(), nr)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err)
		return
	}
	writeJSON(w, http.StatusCreated, res)
}

func (h *Reservation) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.Store.QueryByID(r.Context(), id)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err)
		return
	}
	writeJSON(w, http.StatusCreated, res)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, code int, err error) {
	writeJSON(w, code, map[string]string{"error": err.Error()})
}
