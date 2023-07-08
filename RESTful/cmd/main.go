package main

import (
	"RESTful/pkg/constants"
	"RESTful/pkg/sqlite"
	"RESTful/router"
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load(constants.EnvironmentDirectory)
	db, err := sqlite.NewSQLiteDB()

	if err != nil {
		log.Fatal(err)
	}
	db, err = sqlite.Migrate()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	router.SetupRoutes()

	log.Println("Server started on http://localhost:8080")
	err = http.ListenAndServe(":8050", nil)
	if err != nil {
		return
	}
}
