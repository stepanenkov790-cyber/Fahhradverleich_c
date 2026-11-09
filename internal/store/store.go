package store

import (
	"errors"
	"sync"

	"fahrradverleih/internal/domain"
)

type Store struct {
	mu                                                  sync.RWMutex
	bikes                                               []domain.Bike
	customers                                           []domain.Customer
	rentals                                             []domain.Rental
	reservations                                        []domain.Reservation
	nextBike, nextCustomer, nextRental, nextReservation int
}

func New() *Store {
	return &Store{
		bikes: []domain.Bike{
			{ID: 1, Name: "Trekking Bike", Type: "Trekking", Size: "M", DailyPrice: 18, Status: "available"},
			{ID: 2, Name: "Mountainbike Pro", Type: "Mountainbike", Size: "L", DailyPrice: 25, Status: "rented"},
			{ID: 3, Name: "City Bike", Type: "Citybike", Size: "M", DailyPrice: 14, Status: "available"},
		},
		customers: []domain.Customer{
			{ID: 1, Name: "Max Mustermann", Email: "max@example.com"},
			{ID: 2, Name: "Anna Beispiel", Email: "anna@example.com"},
		},
		rentals:      []domain.Rental{{ID: 1, BikeID: 2, CustomerID: 1, Date: "2026-11-05", Days: 1, DailyPrice: 25, TotalPrice: 25}},
		reservations: []domain.Reservation{},
		nextBike:     4, nextCustomer: 3, nextRental: 2, nextReservation: 1,
	}
}

func (s *Store) Dashboard() domain.Dashboard {
	s.mu.RLock()
	defer s.mu.RUnlock()
	d := domain.Dashboard{TotalBikes: len(s.bikes), Customers: len(s.customers)}
	for _, b := range s.bikes {
		if b.Status == "available" {
			d.AvailableBikes++
		} else {
			d.RentedBikes++
		}
	}
	return d
}
func (s *Store) Bikes() []domain.Bike {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.Bike{}, s.bikes...)
}
func (s *Store) AddBike(v domain.Bike) (domain.Bike, error) {
	if v.DailyPrice <= 0 {
		return v, errors.New("Tagespreis muss größer als 0 sein")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	v.ID = s.nextBike
	s.nextBike++
	if v.Status == "" {
		v.Status = "available"
	}
	s.bikes = append(s.bikes, v)
	return v, nil
}
func (s *Store) DeleteBike(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, r := range s.rentals {
		if r.BikeID == id {
			return errors.New("Fahrrad kann nicht gelöscht werden: Es gibt einen Verleihdatensatz")
		}
	}
	for _, r := range s.reservations {
		if r.BikeID == id {
			return errors.New("Fahrrad kann nicht gelöscht werden: Es gibt eine Reservierung")
		}
	}
	for i, b := range s.bikes {
		if b.ID == id {
			s.bikes = append(s.bikes[:i], s.bikes[i+1:]...)
			return nil
		}
	}
	return errors.New("Fahrrad nicht gefunden")
}
func (s *Store) Customers() []domain.Customer {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.Customer{}, s.customers...)
}
func (s *Store) AddCustomer(v domain.Customer) domain.Customer {
	s.mu.Lock()
	defer s.mu.Unlock()
	v.ID = s.nextCustomer
	s.nextCustomer++
	s.customers = append(s.customers, v)
	return v
}
func (s *Store) DeleteCustomer(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, r := range s.rentals {
		if r.CustomerID == id {
			return errors.New("Kunde hat einen Verleihdatensatz und kann nicht gelöscht werden")
		}
	}
	for _, r := range s.reservations {
		if r.CustomerID == id {
			return errors.New("Kunde hat eine Reservierung und kann nicht gelöscht werden")
		}
	}
	for i, c := range s.customers {
		if c.ID == id {
			s.customers = append(s.customers[:i], s.customers[i+1:]...)
			return nil
		}
	}
	return errors.New("Kunde nicht gefunden")
}
func (s *Store) AddRental(v domain.Rental) (domain.Rental, error) {
	if v.Days < 1 || v.Days > 365 {
		return v, errors.New("Mietdauer muss zwischen 1 und 365 Tagen liegen")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	bi := -1
	for i, b := range s.bikes {
		if b.ID == v.BikeID && b.Status == "available" {
			bi = i
			break
		}
	}
	if bi < 0 {
		return v, errors.New("Fahrrad nicht verfügbar oder nicht gefunden")
	}
	cf := false
	for _, c := range s.customers {
		if c.ID == v.CustomerID {
			cf = true
			break
		}
	}
	if !cf {
		return v, errors.New("Kunde nicht gefunden")
	}
	v.ID = s.nextRental
	s.nextRental++
	v.DailyPrice = s.bikes[bi].DailyPrice
	v.TotalPrice = v.DailyPrice * float64(v.Days)
	s.bikes[bi].Status = "rented"
	s.rentals = append(s.rentals, v)
	return v, nil
}
func (s *Store) Rentals() []domain.Rental {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.Rental{}, s.rentals...)
}
func (s *Store) AddReservation(v domain.Reservation) (domain.Reservation, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	bf, cf := false, false
	for _, b := range s.bikes {
		if b.ID == v.BikeID {
			bf = true
		}
	}
	for _, c := range s.customers {
		if c.ID == v.CustomerID {
			cf = true
		}
	}
	if !bf || !cf {
		return v, errors.New("Fahrrad oder Kunde nicht gefunden")
	}
	v.ID = s.nextReservation
	s.nextReservation++
	if v.Status == "" {
		v.Status = "offen"
	}
	s.reservations = append(s.reservations, v)
	return v, nil
}
func (s *Store) Reservations() []domain.Reservation {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.Reservation{}, s.reservations...)
}
