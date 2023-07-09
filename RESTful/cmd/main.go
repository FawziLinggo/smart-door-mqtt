package main

import (
	"RESTful/pkg/constants"
	"RESTful/pkg/sqlite"
	"RESTful/router"
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
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
	// health check
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		write, err := w.Write([]byte("Im OK"))
		if err != nil {
			return
		}
		log.Println(write)
	})
	// static file
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started on http://localhost:" + os.Getenv(constants.EnvirontmentPort))
	err = http.ListenAndServe(":"+os.Getenv(constants.EnvirontmentPort), nil)
	if err != nil {
		return
	}
}
