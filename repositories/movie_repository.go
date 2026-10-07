package repositories

import (
	"cinema-booking/models"
	"database/sql"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) GetMovies() ([]models.Movie, error) {
	query := `
	SELECT id,title,duration
	FROM movies
	ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Duration,
		)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MovieRepository) AddMovie(movieRecord *models.Movie) error {

	query := `
	INSERT INTO movies (title, duration)
	VALUES (?,?)
	`

	_, err := r.db.Exec(
		query,
		movieRecord.Title,
		movieRecord.Duration,
	)
	if err != nil {
		return err
	}

	return err
}

func (r *MovieRepository) GetMovieByID(id uint) (*models.Movie, error) {
	query := `
	SELECT id,title,duration
	FROM movies
	WHERE id = ?
	`

	var movie models.Movie
	err := r.db.QueryRow(query, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Duration,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}
