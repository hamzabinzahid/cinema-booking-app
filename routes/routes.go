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
) {
	// Movie routes
	http.HandleFunc("GET /api/movies", movieController.GetMovies)
	http.HandleFunc("GET /api/movies/{id}", movieController.GetMovieByID)
	http.HandleFunc("POST /api/movies", movieController.AddMovie)

	// Cinema routes
	http.HandleFunc("GET /api/cinemas", cinemaController.GetCinemas)

	// Showtime routes
	http.HandleFunc("GET /api/showtimes", showtimeController.GetShowtimes)
	http.HandleFunc("GET /api/showtimes/{id}", showtimeController.GetShowtimeByID)

	// Booking routes
	http.HandleFunc("POST /api/bookings", bookingController.CreateBooking)
}
