package main

import (
	"log"
	"net/http"

	"cinema-booking/controllers"
	"cinema-booking/repositories"
	"cinema-booking/routes"
	"cinema-booking/services"
)

func main() {

	// Database connection will be added here.
	// db, err := database.Connect(...)

	// Repositories
	movieRepository := repositories.NewMovieRepository(nil)
	showtimeRepository := repositories.NewShowtimeRepository(nil)
	bookingRepository := repositories.NewBookingRepository(nil)

	// Services
	movieService := services.NewMovieService(movieRepository)
	showtimeService := services.NewShowtimeService(showtimeRepository)
	bookingService := services.NewBookingService(bookingRepository)

	// Controllers
	movieController := controllers.NewMovieController(movieService)
	showtimeController := controllers.NewShowtimeController(showtimeService)
	bookingController := controllers.NewBookingController(bookingService)

	// Routes
	routes.RegisterRoutes(
		movieController,
		showtimeController,
		bookingController,
	)

	log.Println("Cinema server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
