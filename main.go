package main

import (
	"log"
	"net/http"

	"cinema-booking/controllers"
	"cinema-booking/database"
	"cinema-booking/repositories"
	"cinema-booking/routes"
	"cinema-booking/services"
)

func main() {

	// Database connection will be added here.
	db, err := database.Connect(
		"localhost",
		"3306",
		"root",
		"",
		"cinema_booking",
	)

	if err != nil {
		log.Fatal("Database connection failed with error : ", err)
	}
	defer db.Close()
	log.Println("Database connected successfully")

	// Repositories
	movieRepository := repositories.NewMovieRepository(db)
	showtimeRepository := repositories.NewShowtimeRepository(db)
	bookingRepository := repositories.NewBookingRepository(db)
	cinemaRepository := repositories.NewCinemaRepository(db)

	// Services
	movieService := services.NewMovieService(movieRepository)
	showtimeService := services.NewShowtimeService(showtimeRepository)
	bookingService := services.NewBookingService(bookingRepository)
	cinemaService := services.NewCinemaService(cinemaRepository)

	// Controllers
	movieController := controllers.NewMovieController(movieService)
	showtimeController := controllers.NewShowtimeController(showtimeService)
	bookingController := controllers.NewBookingController(bookingService)
	cinemaController := controllers.NewCinemaController(cinemaService)

	// Routes
	routes.RegisterRoutes(
		movieController,
		showtimeController,
		bookingController,
	)

	log.Println("Cinema server running on http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
