package helpers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"os"
)

func SavePathToTable(pathDB string) error {
	db, err := sql.Open("sqlite3", os.Getenv("DATABASE_PATH"))
	if err != nil {
		return err
	}

	pathDB = "static/img/" + pathDB

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	query := "INSERT INTO images (id, path) VALUES (?,?)"
	_, err = db.Exec(query, id.String(), pathDB)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			return
		}
	}(db)
	log.Println("Save path to table success")

	return nil
}
