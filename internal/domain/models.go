package domain

type Bike struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Size       string  `json:"size"`
	DailyPrice float64 `json:"dailyPrice"`
	Status     string  `json:"status"` // available oder rented
}

type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Rental struct {
	ID         int     `json:"id"`
	BikeID     int     `json:"bikeId"`
	CustomerID int     `json:"customerId"`
	Date       string  `json:"date"`
	Days       int     `json:"days"`
	DailyPrice float64 `json:"dailyPrice"`
	TotalPrice float64 `json:"totalPrice"`
}

type Reservation struct {
	ID         int    `json:"id"`
	BikeID     int    `json:"bikeId"`
	CustomerID int    `json:"customerId"`
	Date       string `json:"date"`
	Status     string `json:"status"`
}

type Dashboard struct {
	AvailableBikes int `json:"availableBikes"`
	RentedBikes    int `json:"rentedBikes"`
	TotalBikes     int `json:"totalBikes"`
	Customers      int `json:"customers"`
}
