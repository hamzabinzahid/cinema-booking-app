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
	SELECT
	cinemas.id,
	cinemas.name,
	screens.id,
	screens.name,
	screens.cinema_id
	FROM cinemas
	LEFT JOIN screens ON cinemas.id=screens.cinema_id
	ORDER BY cinemas.id, screens.id
	`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	var cinemas []models.Cinema
	cinemaMap := make(map[uint]uint)
	defer rows.Close()
	for rows.Next() {
		var cinemaID uint
		var cinemaName string
		var screenID sql.NullInt64
		var screenName sql.NullString
		var screenCinemaID sql.NullInt64

		err := rows.Scan(
			&cinemaID,
			&cinemaName,
			&screenID,
			&screenName,
			&screenCinemaID,
		)

		if err != nil {
			return nil, err
		}

		index, exists := cinemaMap[cinemaID]
		if !exists {
			cinemas = append(cinemas, models.Cinema{
				CinemaDetails: models.CinemaDetails{
					ID:   cinemaID,
					Name: cinemaName,
				},
				Screens: []models.ScreenSingle{},
			})
			index = uint(len(cinemas) - 1)
			cinemaMap[cinemaID] = index
		}

		if screenID.Valid {
			cinemas[index].Screens = append(cinemas[index].Screens, models.ScreenSingle{
				ID:       uint(screenID.Int64),
				Name:     screenName.String,
				CinemaID: uint(cinemaID),
			})
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return cinemas, nil
}

func (c *CinemaRepository) GetCinemaById(id int) (*models.Cinema, error) {
	query := `
	SELECT
	cinemas.id,
	cinemas.name,
	screens.id,
	screens.name,
	screens.cinema_id
	FROM cinemas
	LEFT JOIN screens ON cinemas.id=screens.cinema_id
	WHERE cinemas.id=?
	`

	rows, err := c.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var cinema models.Cinema
	cinema.Screens = []models.ScreenSingle{}
	found := false
	for rows.Next() {
		var CinemaID uint
		var CinemaName string
		var ScreenID sql.NullInt64
		var ScreenName sql.NullString
		var screenCinemaID sql.NullInt64

		err := rows.Scan(
			&CinemaID,
			&CinemaName,
			&ScreenID,
			&ScreenName,
			&screenCinemaID,
		)

		if err != nil {
			return nil, err
		}

		if !found {
			cinema.ID = CinemaID
			cinema.Name = CinemaName
			found = true
		}

		if ScreenID.Valid {
			cinema.Screens = append(cinema.Screens, models.ScreenSingle{
				ID:       uint(ScreenID.Int64),
				Name:     ScreenName.String,
				CinemaID: uint(screenCinemaID.Int64),
			})
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &cinema, nil
}

func (c *CinemaRepository) AddCinema(cinemaDetails models.CinemaDetails) (*models.CinemaDetails, error) {
	query := `
	INSERT INTO cinemas(id,name)
	VALUES(?,?)
	`

	data, err := c.db.Exec(query, cinemaDetails.ID, cinemaDetails.Name)
	if err != nil {
		return nil, err
	}

	// Get Id of Last Added Record
	lastId, err := data.LastInsertId()
	if err != nil {
		return nil, err
	}

	query = `
	SELECT
	cinemas.id,
	cinemas.name
	FROM cinemas
	WHERE cinemas.id=?
	`

	latestRecord := c.db.QueryRow(query, lastId)
	var latestCinema models.CinemaDetails
	latestRecord.Scan(
		&latestCinema.ID,
		&latestCinema.Name,
	)

	return &latestCinema, nil
}
