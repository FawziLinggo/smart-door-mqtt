package sqlite

import (
	"RESTful/app/repositories"
	"RESTful/pkg/constants"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func NewSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", os.Getenv(constants.EnvirontmentDatabasePath))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", os.Getenv(constants.EnvirontmentDatabasePath))
	if err != nil {
		return nil, err
	}
	migration := repositories.NewUserRepository(db)
	err = migration.Migrate()
	if err != nil {
		log.Fatal("Error migration: ", err)
	}
	return db, nil
}
