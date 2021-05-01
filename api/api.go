package api

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	//"github.com/Alexandr59/golang-training-Theater/pkg/data"
	"golang-training-Theater/pkg/data"
)

func ServerTheaterResource(r *mux.Router, conn *gorm.DB) {

	account := &accountAPI{data: data.NewAccountData(conn)}
	genre := &genreAPI{data: data.NewGenreData(conn)}
	hall := &hallAPI{data: data.NewHallData(conn)}
	location := &locationAPI{data: data.NewLocationData(conn)}
	performance := &performanceAPI{data: data.NewPerformanceData(conn)}
	place := &placeAPI{data: data.NewPlaceData(conn)}
	poster := &posterAPI{data: data.NewPosterData(conn)}
	price := &priceAPI{data: data.NewPriceData(conn)}
	role := &roleAPI{data: data.NewRoleData(conn)}
	schedule := &scheduleAPI{data: data.NewScheduleData(conn)}
	sector := &sectorAPI{data: data.NewSectorData(conn)}
	ticket := &ticketAPI{data: data.NewTicketData(conn)}
	user := &userAPI{data: data.NewUserData(conn)}

	r.HandleFunc("/account", account.createAccount).Methods("POST")
	r.HandleFunc("/genre", genre.createGenre).Methods("POST")
	r.HandleFunc("/hall", hall.createHall).Methods("POST")
	r.HandleFunc("/location", location.createLocation).Methods("POST")
	r.HandleFunc("/performance", performance.createPerformance).Methods("POST")
	r.HandleFunc("/place", place.createPlace).Methods("POST")
	r.HandleFunc("/poster", poster.createPoster).Methods("POST")
	r.HandleFunc("/price", price.createPrice).Methods("POST")
	r.HandleFunc("/role", role.createRole).Methods("POST")
	r.HandleFunc("/schedule", schedule.createSchedule).Methods("POST")
	r.HandleFunc("/sector", sector.createSector).Methods("POST")
	r.HandleFunc("/ticket", ticket.createTicket).Methods("POST")
	r.HandleFunc("/user", user.createUser).Methods("POST")

	r.HandleFunc("/account", account.deleteAccountById).Methods("DELETE")
	r.HandleFunc("/genre", genre.deleteGenreById).Methods("DELETE")
	r.HandleFunc("/hall", hall.deleteHallById).Methods("DELETE")
	r.HandleFunc("/location", location.deleteLocationById).Methods("DELETE")
	r.HandleFunc("/performance", performance.deletePerformanceById).Methods("DELETE")
	r.HandleFunc("/place", place.deletePlaceById).Methods("DELETE")
	r.HandleFunc("/poster", poster.deletePosterById).Methods("DELETE")
	r.HandleFunc("/price", price.deletePriceById).Methods("DELETE")
	r.HandleFunc("/role", role.deleteRoleById).Methods("DELETE")
	r.HandleFunc("/schedule", schedule.deleteScheduleById).Methods("DELETE")
	r.HandleFunc("/sector", sector.deleteSectorById).Methods("DELETE")
	r.HandleFunc("/ticket", ticket.deleteTicketById).Methods("DELETE")
	r.HandleFunc("/user", user.deleteUserById).Methods("DELETE")

	r.HandleFunc("/account", account.updateAccount).Methods("PUT")
	r.HandleFunc("/genre", genre.updateGenre).Methods("PUT")
	r.HandleFunc("/hall", hall.updateHall).Methods("PUT")
	r.HandleFunc("/location", location.updateLocation).Methods("PUT")
	r.HandleFunc("/performance", performance.updatePerformance).Methods("PUT")
	r.HandleFunc("/place", place.updatePlace).Methods("PUT")
	r.HandleFunc("/poster", poster.updatePoster).Methods("PUT")
	r.HandleFunc("/price", price.updatePrice).Methods("PUT")
	r.HandleFunc("/role", role.updateRole).Methods("PUT")
	r.HandleFunc("/schedule", schedule.updateSchedule).Methods("PUT")
	r.HandleFunc("/sector", sector.updateSector).Methods("PUT")
	r.HandleFunc("/ticket", ticket.updateTicket).Methods("PUT")
	r.HandleFunc("/user", user.updateUser).Methods("PUT")

	r.HandleFunc("/account", account.getAccountById).Methods("GET")
	r.HandleFunc("/genre", genre.getGenreById).Methods("GET")
	r.HandleFunc("/hall", hall.getHallById).Methods("GET")
	r.HandleFunc("/location", location.getLocationById).Methods("GET")
	r.HandleFunc("/performance", performance.getPerformanceById).Methods("GET")
	r.HandleFunc("/place", place.getPlaceById).Methods("GET")
	r.HandleFunc("/poster", poster.getPosterById).Methods("GET")
	r.HandleFunc("/price", price.getPriceById).Methods("GET")
	r.HandleFunc("/role", role.getRoleById).Methods("GET")
	r.HandleFunc("/schedule", schedule.getScheduleById).Methods("GET")
	r.HandleFunc("/sector", sector.getSectorById).Methods("GET")
	r.HandleFunc("/ticket", ticket.getTicketById).Methods("GET")
	r.HandleFunc("/user", user.getUserById).Methods("GET")

	r.HandleFunc("/tickets", ticket.getAllTickets).Methods("GET")
	r.HandleFunc("/posters", poster.getAllPosters).Methods("GET")
	r.HandleFunc("/users", user.getAllUsers).Methods("GET")
}
