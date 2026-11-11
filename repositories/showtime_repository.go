package repositories

import (
	"cinema-booking/models"
	"database/sql"
)

type ShowtimeRepository struct {
	db *sql.DB
}

func NewShowtimeRepository(db *sql.DB) *ShowtimeRepository {
	return &ShowtimeRepository{
		db: db,
	}
}

func (r *ShowtimeRepository) GetShowtimes() ([]models.ShowTimeRel, error) {
	query := `
	SELECT
		showtimes.id,
		showtimes.movie_id,
		showtimes.screen_id,
		showtimes.start_time,

		movies.id AS movie_id,
		movies.title,
		movies.duration,

		screens.id AS screen_id,
		screens.name,

		cinemas.id AS cinema_id,
		cinemas.name AS cinema_name
	FROM showtimes
	JOIN movies ON showtimes.movie_id=movies.id
	JOIN screens ON showtimes.screen_id=screens.id
	JOIN cinemas ON screens.cinema_id=cinemas.id
	ORDER BY showtimes.id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var showTimes []models.ShowTimeRel
	for rows.Next() {
		var showTime models.ShowTimeRel
		err := rows.Scan(
			&showTime.ID,
			&showTime.MovieID,
			&showTime.ScreenID,
			&showTime.StartTime,
			&showTime.Movie.ID,
			&showTime.Movie.Title,
			&showTime.Movie.Duration,
			&showTime.Screen.ID,
			&showTime.Screen.Name,
			&showTime.Screen.CinemaDetails.ID,
			&showTime.Screen.CinemaDetails.Name,
		)

		if err != nil {
			return nil, err
		}

		showTimes = append(showTimes, showTime)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return showTimes, nil
}

func (r *ShowtimeRepository) GetShowtimesByMovie(movieId uint) ([]models.ShowTimeRel, error) {
	query := `
	SELECT
		showtimes.id,
		showtimes.movie_id,
		showtimes.screen_id,
		showtimes.start_time,

		movies.id AS movie_id,
		movies.title,
		movies.duration,

		screens.id AS screen_id,
		screens.name,

		cinemas.id AS cinema_id,
		cinemas.name AS cinema_name
	FROM showtimes
	JOIN movies ON showtimes.movie_id=movies.id
	JOIN screens ON showtimes.screen_id=screens.id
	JOIN cinemas ON screens.cinema_id=cinemas.id
	WHERE showtimes.movie_id=?
	ORDER BY showtimes.id
	`

	rows, err := r.db.Query(query, movieId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var showTimes []models.ShowTimeRel
	for rows.Next() {
		var showTime models.ShowTimeRel
		err := rows.Scan(
			&showTime.ID,
			&showTime.MovieID,
			&showTime.ScreenID,
			&showTime.StartTime,
			&showTime.Movie.ID,
			&showTime.Movie.Title,
			&showTime.Movie.Duration,
			&showTime.Screen.ID,
			&showTime.Screen.Name,
			&showTime.Screen.CinemaDetails.ID,
			&showTime.Screen.CinemaDetails.Name,
		)

		if err != nil {
			return nil, err
		}

		showTimes = append(showTimes, showTime)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return showTimes, nil
}

func (r *ShowtimeRepository) GetShowtimesById(showTimeId uint) (*models.ShowTimeRel, error) {
	query := `
	SELECT
		showtimes.id,
		showtimes.movie_id,
		showtimes.screen_id,
		showtimes.start_time,

		movies.id AS movie_id,
		movies.title,
		movies.duration,

		screens.id AS screen_id,
		screens.name,

		cinemas.id AS cinema_id,
		cinemas.name AS cinema_name
	FROM showtimes
	JOIN movies ON showtimes.movie_id=movies.id
	JOIN screens ON showtimes.screen_id=screens.id
	JOIN cinemas ON screens.cinema_id=cinemas.id
	WHERE showtimes.id=?
	ORDER BY showtimes.id
	`

	var showTime models.ShowTimeRel
	err := r.db.QueryRow(query, showTimeId).Scan(
		&showTime.ID,
		&showTime.MovieID,
		&showTime.ScreenID,
		&showTime.StartTime,
		&showTime.Movie.ID,
		&showTime.Movie.Title,
		&showTime.Movie.Duration,
		&showTime.Screen.ID,
		&showTime.Screen.Name,
		&showTime.Screen.CinemaDetails.ID,
		&showTime.Screen.CinemaDetails.Name,
	)
	if err != nil {
		return nil, err
	}

	return &showTime, nil
}

func (r *ShowtimeRepository) AddShowtime(showTime models.Showtime) (*models.ShowTimeRel, error) {
	query := `
	INSERT INTO showtimes(movie_id,screen_id,start_time)
	VALUES(?,?,?)
	`

	data, err := r.db.Exec(query, showTime.MovieID, showTime.ScreenID, showTime.StartTime)
	if err != nil {
		return nil, err
	}

	LastInsertId, err := data.LastInsertId()
	if err != nil {
		return nil, err
	}

	query = `
	SELECT
		showtimes.id,
		showtimes.movie_id,
		showtimes.screen_id,
		showtimes.start_time,

		movies.id AS movie_id,
		movies.title,
		movies.duration,

		screens.id AS screen_id,
		screens.name,

		cinemas.id AS cinema_id,
		cinemas.name AS cinema_name
	FROM showtimes
	JOIN movies ON showtimes.movie_id=movies.id
	JOIN screens ON showtimes.screen_id=screens.id
	JOIN cinemas ON screens.cinema_id=cinemas.id
	WHERE showtimes.id=?
	ORDER BY showtimes.id
	`

	var lastRecord models.ShowTimeRel
	err = r.db.QueryRow(query, LastInsertId).Scan(
		&lastRecord.ID,
		&lastRecord.MovieID,
		&lastRecord.ScreenID,
		&lastRecord.StartTime,
		&lastRecord.Movie.ID,
		&lastRecord.Movie.Title,
		&lastRecord.Movie.Duration,
		&lastRecord.Screen.ID,
		&lastRecord.Screen.Name,
		&lastRecord.Screen.CinemaDetails.ID,
		&lastRecord.Screen.CinemaDetails.Name,
	)

	if err != nil {
		return nil, err
	}

	return &lastRecord, nil

}
