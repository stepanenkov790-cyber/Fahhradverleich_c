package main

import (
	"log"
	"net/http"

	"fahrradverleih/internal/handler"
	"fahrradverleih/internal/store"
)

func main() {
	s := store.New()
	h := handler.New(s)
	mux := http.NewServeMux()
	mux.HandleFunc("/api/dashboard", h.Dashboard)
	mux.HandleFunc("/api/bikes", h.Bikes)
	mux.HandleFunc("/api/bikes/", h.BikeByID)
	mux.HandleFunc("/api/customers", h.Customers)
	mux.HandleFunc("/api/customers/", h.CustomerByID)
	mux.HandleFunc("/api/rentals", h.Rentals)
	mux.HandleFunc("/api/reservations", h.Reservations)
	mux.Handle("/", http.FileServer(http.Dir("./web")))
	log.Println("Fahrradverleih läuft auf http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
