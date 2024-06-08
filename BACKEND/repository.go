package main

import (
	"fmt"
)

func RegisterUser(user User) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	return err
}

func AuthenticateUser(user User) (string, error) {
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username=$1", user.Username).Scan(&storedPassword)
	if err != nil {
		return "", err
	}
	if !CheckPasswordHash(user.Password, storedPassword) {
		return "", fmt.Errorf("invalid credentials")
	}
	// Implement JWT token generation here
	token, err := GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func CreateReservation(reservation Reservation) error {
	_, err := db.Exec("INSERT INTO reservations (user_id, car_id, extras) VALUES ($1, $2, $3)", reservation.UserID, reservation.CarID, reservation.Extras)
	return err
}

func GetReservations(userID int) ([]Reservation, error) {
	rows, err := db.Query("SELECT id, user_id, car_id, extras FROM reservations WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reservations []Reservation
	for rows.Next() {
		var reservation Reservation
		if err := rows.Scan(&reservation.ID, &reservation.UserID, &reservation.CarID, &reservation.Extras); err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}

func DeleteReservation(id int) error {
	_, err := db.Exec("DELETE FROM reservations WHERE id=$1", id)
	return err
}

func GenerateReport(userID int) ([]Reservation, error) {
	return GetReservations(userID)
}
