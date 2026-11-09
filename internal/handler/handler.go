package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"fahrradverleih/internal/domain"
	"fahrradverleih/internal/store"
)

type Handler struct{ store *store.Store }

func New(s *store.Store) *Handler { return &Handler{store: s} }
func write(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func decode(r *http.Request, v any) error { return json.NewDecoder(r.Body).Decode(v) }
func parseID(path, prefix string) (int, error) {
	return strconv.Atoi(strings.Trim(strings.TrimPrefix(path, prefix), "/"))
}
func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Methode nicht erlaubt", 405)
		return
	}
	write(w, 200, h.store.Dashboard())
}
func (h *Handler) Bikes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		write(w, 200, h.store.Bikes())
	case "POST":
		var v domain.Bike
		if decode(r, &v) != nil {
			http.Error(w, "Ungültige Daten", 400)
			return
		}
		x, e := h.store.AddBike(v)
		if e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		write(w, 201, x)
	default:
		http.Error(w, "Methode nicht erlaubt", 405)
	}
}
func (h *Handler) BikeByID(w http.ResponseWriter, r *http.Request) {
	id, e := parseID(r.URL.Path, "/api/bikes/")
	if e != nil {
		http.Error(w, "Ungültige ID", 400)
		return
	}
	if r.Method == "DELETE" {
		if e = h.store.DeleteBike(id); e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		w.WriteHeader(204)
		return
	}
	http.Error(w, "Methode nicht erlaubt", 405)
}
func (h *Handler) Customers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		write(w, 200, h.store.Customers())
	case "POST":
		var v domain.Customer
		if decode(r, &v) != nil {
			http.Error(w, "Ungültige Daten", 400)
			return
		}
		write(w, 201, h.store.AddCustomer(v))
	default:
		http.Error(w, "Methode nicht erlaubt", 405)
	}
}
func (h *Handler) CustomerByID(w http.ResponseWriter, r *http.Request) {
	id, e := parseID(r.URL.Path, "/api/customers/")
	if e != nil {
		http.Error(w, "Ungültige ID", 400)
		return
	}
	if r.Method == "DELETE" {
		if e = h.store.DeleteCustomer(id); e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		w.WriteHeader(204)
		return
	}
	http.Error(w, "Methode nicht erlaubt", 405)
}
func (h *Handler) Rentals(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		write(w, 200, h.store.Rentals())
	case "POST":
		var v domain.Rental
		if decode(r, &v) != nil {
			http.Error(w, "Ungültige Daten", 400)
			return
		}
		x, e := h.store.AddRental(v)
		if e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		write(w, 201, x)
	default:
		http.Error(w, "Methode nicht erlaubt", 405)
	}
}
func (h *Handler) Reservations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		write(w, 200, h.store.Reservations())
	case "POST":
		var v domain.Reservation
		if decode(r, &v) != nil {
			http.Error(w, "Ungültige Daten", 400)
			return
		}
		x, e := h.store.AddReservation(v)
		if e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		write(w, 201, x)
	default:
		http.Error(w, "Methode nicht erlaubt", 405)
	}
}
