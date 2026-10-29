package repositories

import (
	"cinema-booking/models"
	"database/sql"
)

type ScreenRepository struct {
	db *sql.DB
}

func NewScreenRepository(db *sql.DB) *ScreenRepository {
	return &ScreenRepository{
		db: db,
	}
}

func (s *ScreenRepository) AddScreen(screen models.Screen) (*models.Screen, error) {
	query := `
	INSERT INTO screens (name,cinema_id)
	VALUES (?,?)
	`

	result, err := s.db.Exec(
		query,
		screen.Name,
		screen.CinemaID,
	)

	if err != nil {
		return nil, err
	}

	screenId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	query = `
	SELECT
		screens.id,
		screens.name,
		screens.cinema_id,
		cinemas.id,
		cinemas.name
	FROM screens
	JOIN cinemas ON screens.cinema_id=cinemas.id
	WHERE screens.id=?
	`
	var singleScreen models.Screen
	err = s.db.QueryRow(query, screenId).Scan(
		&singleScreen.ID,
		&singleScreen.Name,
		&singleScreen.CinemaID,
		&singleScreen.CinemaDetails.ID,
		&singleScreen.CinemaDetails.Name,
	)

	if err != nil {
		return nil, err
	}

	return &singleScreen, nil
}

func (s *ScreenRepository) GetScreens() ([]models.Screen, error) {
	query := `
	SELECT
		screens.id,
		screens.cinema_id,
		screens.name,
		cinemas.id,
		cinemas.name
	FROM screens
	JOIN cinemas ON screens.cinema_id=cinemas.id
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var screens []models.Screen
	for rows.Next() {
		var singleScreen models.Screen
		err := rows.Scan(
			&singleScreen.ID,
			&singleScreen.CinemaID,
			&singleScreen.Name,
			&singleScreen.CinemaDetails.ID,
			&singleScreen.CinemaDetails.Name,
		)

		if err != nil {
			return nil, err
		}

		screens = append(screens, singleScreen)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return screens, nil
}

func (s *ScreenRepository) GetScreenById(id uint) (*models.Screen, error) {
	query := `SELECT
		screens.id,
		screens.name,
		screens.cinema_id,
		cinemas.id,
		cinemas.name
	FROM screens
	JOIN cinemas ON screens.cinema_id=cinemas.id
	WHERE screens.id=?
	`

	var screenData models.Screen
	err := s.db.QueryRow(query, id).Scan(
		&screenData.ID,
		&screenData.Name,
		&screenData.CinemaID,
		&screenData.CinemaDetails.ID,
		&screenData.CinemaDetails.Name,
	)

	if err != nil {
		return nil, err
	}

	return &screenData, nil
}
