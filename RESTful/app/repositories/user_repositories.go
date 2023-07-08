package repositories

import (
	"RESTful/internal/domain"
	"RESTful/pkg/constants"
	"RESTful/pkg/helpers"
	"database/sql"
	"errors"
	"log"
	"os"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Migrate() error {
	username := os.Getenv(constants.EnvirontmentUsername)
	password := os.Getenv(constants.EnvirontmentPassword)
	passwordHash, err := helpers.HashPassword(password)
	log.Println("username: ", username+" passwordHash : ", passwordHash)

	if err != nil {
		return err
	}

	// Create table users
	query := `
		CREATE TABLE IF NOT EXISTS users (
		    			id INTEGER PRIMARY KEY AUTOINCREMENT,
		    			username TEXT NOT NULL,
		    			password TEXT NOT NULL
		                                 )
    `
	_, err = ur.db.Exec(query)
	if err != nil {
		log.Println("Error create table users: ", err)
		return err
	}
	// Check len rows if == 0 then insert dummy data
	var count int
	err = ur.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Table users already seeded")
		return nil
	}

	// Insert dummy data
	query = "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err = ur.db.Exec(query, username, passwordHash)
	if err != nil {
		return err
	}

	log.Println("Seed Table users done")

	return nil

}

func (ur *UserRepository) GetByName(name string) (*domain.User, error) {
	row := ur.db.QueryRow("SELECT * FROM users WHERE username = ?", name)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(constants.ErrNotExists)
		}
		return nil, err
	}

	return &user, nil
}
