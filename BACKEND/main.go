package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		claims, err := ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		r.Header.Set("UserID", claims.Username)
		next.ServeHTTP(w, r)
	})
}

// ConectarDB crea una conexión a la base de datos PostgreSQL
func ConectarDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	// Obteniendo las variables de entorno para la conexión a la base de datos
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Creando un objeto de conexión a PostgreSQL
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := ConectarDB(connectionString)
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}
	defer db.Close()

	// Verificando la conexión a la base de datos
	err = db.Ping()
	if err != nil {
		log.Fatalln("error haciendo ping a la base de datos", err.Error())
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")

	protectedRoutes := r.PathPrefix("/").Subrouter()
	protectedRoutes.Use(authMiddleware)

	protectedRoutes.HandleFunc("/reservations", CreateReservationHandler).Methods("POST")
	protectedRoutes.HandleFunc("/reservations", GetReservationsHandler).Methods("GET")
	protectedRoutes.HandleFunc("/reservations/{id}", DeleteReservationHandler).Methods("DELETE")
	protectedRoutes.HandleFunc("/report", ReportHandler).Methods("GET")

	log.Println("Server running on port 8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
