package main

import (
	"database/sql"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func init() {
	var err error
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	db, err = sql.Open("postgres",
		"user="+dbUser+" password="+dbPassword+" dbname="+dbName+" sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash checks if the provided password matches the hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
