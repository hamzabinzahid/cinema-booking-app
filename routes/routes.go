package routes

import (
	"net/http"

	"cinema-booking/controllers"
)

func RegisterRoutes(
	movieController *controllers.MovieController,
	showtimeController *controllers.ShowtimeController,
	bookingController *controllers.BookingController,
	cinemaController *controllers.CinemaController,
	screenController *controllers.ScreenController,
) {
	// Movie routes
	http.HandleFunc("GET /api/movies", movieController.GetMovies)
	http.HandleFunc("GET /api/movies/{id}", movieController.GetMovieByID)
	http.HandleFunc("POST /api/add-movie", movieController.AddMovie)

	// Screen routes
	http.HandleFunc("GET /api/screens", screenController.GetScreens)
	http.HandleFunc("GET /api/screens/{id}", screenController.GetScreenById)
	http.HandleFunc("POST /api/add-screen", screenController.AddScreen)

	// Cinema routes
	http.HandleFunc("GET /api/cinemas", cinemaController.GetCinemas)
	http.HandleFunc("GET /api/cinemas/{id}", cinemaController.GetCinemaById)
	http.HandleFunc("POST /api/add-cinema", cinemaController.AddCinema)

	// Showtime routes
	http.HandleFunc("GET /api/showtimes", showtimeController.GetShowtimes)
	http.HandleFunc("GET /api/showtimes/{id}", showtimeController.GetShowtimeByID)

	// Booking routes
	http.HandleFunc("POST /api/bookings", bookingController.CreateBooking)
}
