package repositories

import "database/sql"

type ImageRepository struct {
	db *sql.DB
}

func NewImageRepository(db *sql.DB) *ImageRepository {
	return &ImageRepository{
		db: db,
	}
}

func (ir *ImageRepository) GetImagePathById(id int) (string, error) {
	var image string
	query := "SELECT path FROM images WHERE id = ?"
	err := ir.db.QueryRow(query, id).Scan(&image)
	if err != nil {
		return "", err
	}
	return image, nil
}

func (ir *ImageRepository) GetAllImagePath() ([]string, error) {
	var paths []string
	query := "SELECT path FROM images"
	rows, err := ir.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	for rows.Next() {
		var image string
		err = rows.Scan(&image)
		if err != nil {
			return nil, err
		}
		paths = append(paths, image)
	}

	return paths, nil
}

func (ir *ImageRepository) DeleteImageById(id int) error {
	query := "DELETE FROM images WHERE id = ?"
	_, err := ir.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
