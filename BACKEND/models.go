package main

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Car struct {
	ID           int    `json:"id"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	FuelType     string `json:"fuel_type"`
	Transmission string `json:"transmission"`
}

type Reservation struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	CarID  int    `json:"car_id"`
	Extras string `json:"extras"`
}
