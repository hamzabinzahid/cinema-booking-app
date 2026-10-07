package repositories

import (
	"cinema-booking/models"
	"database/sql"
)

type CinemaRepository struct {
	db *sql.DB
}

func NewCinemaRepository(db *sql.DB) *CinemaRepository {
	return &CinemaRepository{
		db: db,
	}
}

func (c *CinemaRepository) GetCinemas() ([]models.Cinema, error) {
	query := `
	SELECT id,name
	FROM cinemas
	ORDER BY id
	`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	var cinemas []models.Cinema
	defer rows.Close()
	for rows.Next() {
		var cinema models.Cinema
		err := rows.Scan(
			&cinema.ID,
			&cinema.Name,
		)

		if err != nil {
			return nil, err
		}

		cinemas = append(cinemas, cinema)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return cinemas, nil
}

func (c *CinemaRepository) GetCinemaById(id int) (*models.Cinema, error) {
	query := `
	SELECT id,name
	FROM cinemas
	WHERE id=?
	`

	var cinema models.Cinema
	err := c.db.QueryRow(query, id).Scan(
		&cinema.ID,
		&cinema.Name,
	)

	if err != nil {
		return nil, err
	}

	return &cinema, nil
}

func (c *CinemaRepository) AddCinema(cinemaDetails models.Cinema) (*models.Cinema, error) {
	query := `
	INSERT INTO cinemas(id,name)
	VALUES(?,?)
	`

	_, err := c.db.Exec(query, cinemaDetails.ID, cinemaDetails.Name)
	if err != nil {
		return nil, err
	}

	return &cinemaDetails, nil
}
