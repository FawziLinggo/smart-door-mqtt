package handlers

import (
	"RESTful/app/repositories"
	"RESTful/internal/domain"
	"RESTful/internal/service"
	"RESTful/pkg/constants"
	"RESTful/pkg/helpers"
	"RESTful/pkg/sqlite"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlite.NewSQLiteDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			http.Error(w, "Failed to close database connection", http.StatusInternalServerError)
			return
		}
	}(db)

	userRepository := repositories.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)

	var loginRequest domain.LoginRequest
	if err = json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, "Failed to close request body", http.StatusInternalServerError)
			return
		}
	}(r.Body)

	user, err := authService.Authenticate(loginRequest.Username, loginRequest.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := helpers.HashPassword(os.Getenv(constants.EnvirontmentAPIKey + user.Username))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Login successful",
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
	err = json.NewEncoder(w).Encode(response)

	log.Println("Login successful with username: ", user.Username)
	if err != nil {
		return
	}
}
