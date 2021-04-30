package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Alexandr59/golang-training-Theater/theater-gorm/pkg/data"
	"github.com/gorilla/mux"
)

type theaterAPI struct {
	data *data.TheaterData
}

func ServerTheaterResource(r *mux.Router, data data.TheaterData) {
	api := &theaterAPI{data: &data}
	r.HandleFunc("/tickets", api.getAllTickets).Methods("GET")
	r.HandleFunc("/posters", api.getAllPosters).Methods("GET")
	r.HandleFunc("/users", api.getAllUsers).Methods("GET")

	r.HandleFunc("/account", api.getAccountById).Methods("GET")
	r.HandleFunc("/genre", api.getGenreById).Methods("GET")
	r.HandleFunc("/hall", api.getHallById).Methods("GET")
	r.HandleFunc("/location", api.getLocationById).Methods("GET")
	r.HandleFunc("/performance", api.getPerformanceById).Methods("GET")
	r.HandleFunc("/place", api.getPlaceById).Methods("GET")
	r.HandleFunc("/poster", api.getPosterById).Methods("GET")
	r.HandleFunc("/price", api.getPriceById).Methods("GET")
	r.HandleFunc("/role", api.getRoleById).Methods("GET")
	r.HandleFunc("/schedule", api.getScheduleById).Methods("GET")
	r.HandleFunc("/sector", api.getSectorById).Methods("GET")
	r.HandleFunc("/ticket", api.getTicketById).Methods("GET")
	r.HandleFunc("/user", api.getUserById).Methods("GET")

	r.HandleFunc("/account", api.deleteAccountById).Methods("DELETE")
	r.HandleFunc("/genre", api.deleteGenreById).Methods("DELETE")
	r.HandleFunc("/hall", api.deleteHallById).Methods("DELETE")
	r.HandleFunc("/location", api.deleteLocationById).Methods("DELETE")
	r.HandleFunc("/performance", api.deletePerformanceById).Methods("DELETE")
	r.HandleFunc("/place", api.deletePlaceById).Methods("DELETE")
	r.HandleFunc("/poster", api.deletePosterById).Methods("DELETE")
	r.HandleFunc("/price", api.deletePriceById).Methods("DELETE")
	r.HandleFunc("/role", api.deleteRoleById).Methods("DELETE")
	r.HandleFunc("/schedule", api.deleteScheduleById).Methods("DELETE")
	r.HandleFunc("/sector", api.deleteSectorById).Methods("DELETE")
	r.HandleFunc("/ticket", api.deleteTicketById).Methods("DELETE")
	r.HandleFunc("/user", api.deleteUserById).Methods("DELETE")

	r.HandleFunc("/account", api.createAccount).Methods("POST")
	r.HandleFunc("/genre", api.createGenre).Methods("POST")
	r.HandleFunc("/hall", api.createHall).Methods("POST")
	r.HandleFunc("/location", api.createLocation).Methods("POST")
	r.HandleFunc("/performance", api.createPerformance).Methods("POST")
	r.HandleFunc("/place", api.createPlace).Methods("POST")
	r.HandleFunc("/poster", api.createPoster).Methods("POST")
	r.HandleFunc("/price", api.createPrice).Methods("POST")
	r.HandleFunc("/role", api.createRole).Methods("POST")
	r.HandleFunc("/schedule", api.createSchedule).Methods("POST")
	r.HandleFunc("/sector", api.createSector).Methods("POST")
	r.HandleFunc("/ticket", api.createTicket).Methods("POST")
	r.HandleFunc("/user", api.createUser).Methods("POST")

	r.HandleFunc("/account", api.updateAccount).Methods("PUT")
	r.HandleFunc("/genre", api.updateGenre).Methods("PUT")
	r.HandleFunc("/hall", api.updateHall).Methods("PUT")
	r.HandleFunc("/location", api.updateLocation).Methods("PUT")
	r.HandleFunc("/performance", api.updatePerformance).Methods("PUT")
	r.HandleFunc("/place", api.updatePlace).Methods("PUT")
	r.HandleFunc("/poster", api.updatePoster).Methods("PUT")
	r.HandleFunc("/price", api.updatePrice).Methods("PUT")
	r.HandleFunc("/role", api.updateRole).Methods("PUT")
	r.HandleFunc("/schedule", api.updateSchedule).Methods("PUT")
	r.HandleFunc("/sector", api.updateSector).Methods("PUT")
	r.HandleFunc("/ticket", api.updateTicket).Methods("PUT")
	r.HandleFunc("/user", api.updateUser).Methods("PUT")
}

func (a theaterAPI) getAllTickets(writer http.ResponseWriter, _ *http.Request) {
	tickets, err := a.data.ReadAllTickets()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get tickets"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(tickets)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getAllPosters(writer http.ResponseWriter, _ *http.Request) {
	posters, err := a.data.ReadAllPosters()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get posters"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(posters)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getAllUsers(writer http.ResponseWriter, request *http.Request) {
	account := new(data.Account)

	if n, err := strconv.Atoi(request.URL.Query().Get("idAccount")); err == nil {
		account.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := a.data.ReadAllUsers(*account)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get users"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getAccountById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Account)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdAccount(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get account"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getGenreById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Genre)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdGenre(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get genre"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getHallById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Hall)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdHall(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get hall"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getLocationById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Location)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdLocation(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get location"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getPerformanceById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Performance)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdPerformance(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get performance"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getPlaceById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Place)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdPlace(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get place"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getPosterById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Poster)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdPoster(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get poster"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getPriceById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Price)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdPrice(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get price"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getRoleById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Role)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdRole(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get role"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getScheduleById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Schedule)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdSchedule(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get schedule"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getSectorById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Sector)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdSector(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get sector"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getTicketById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Ticket)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdTicket(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get ticket"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) getUserById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.User)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdUser(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get user"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a theaterAPI) deleteAccountById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Account)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeleteAccount(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete account"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deleteGenreById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Genre)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeleteGenre(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete genre"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deleteHallById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Hall)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeleteHall(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete hall"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deleteLocationById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Location)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeleteLocation(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete location"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deletePerformanceById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Performance)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeletePerformance(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete performance"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deletePlaceById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Place)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeletePlace(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete place"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deletePosterById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Poster)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeletePoster(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete poster"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deletePriceById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Price)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeletePrice(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete price"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deleteRoleById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Role)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeleteRole(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete role"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deleteScheduleById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Schedule)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeleteSchedule(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete schedule"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deleteSectorById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Sector)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeleteSector(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete sector"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deleteTicketById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Ticket)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeleteTicket(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete ticket"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) deleteUserById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.User)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.data.DeleteUser(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete user"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) createAccount(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Account)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddAccount(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create account"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createGenre(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Genre)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddGenre(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create genre"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createHall(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Hall)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddHall(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create hall"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createLocation(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Location)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddLocation(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create location"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createPerformance(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Performance)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddPerformance(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create performance"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createPlace(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Place)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddPlace(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create place"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createPoster(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Poster)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddPoster(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create poster"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createPrice(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Price)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddPrice(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create price"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createRole(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Role)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddRole(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create role"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createSchedule(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Schedule)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddSchedule(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create schedule"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createSector(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Sector)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddSector(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create sector"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createTicket(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Ticket)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddTicket(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create ticket"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) createUser(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.User)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := a.data.AddUser(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create user"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a theaterAPI) updateAccount(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Account)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateAccount(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update account"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updateGenre(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Genre)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateGenre(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update genre"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updateHall(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Hall)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateHall(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update hall"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updateLocation(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Location)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateLocation(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update location"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updatePerformance(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Performance)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdatePerformance(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update performance"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updatePlace(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Place)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdatePlace(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update place"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updatePoster(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Poster)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdatePoster(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update poster"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updatePrice(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Price)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdatePrice(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update price"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updateRole(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Role)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateRole(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update role"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updateSchedule(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Schedule)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateSchedule(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update schedule"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updateSector(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Sector)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateSector(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update sector"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updateTicket(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Ticket)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateTicket(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update ticket"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a theaterAPI) updateUser(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.User)
	err := json.NewDecoder(request.Body).Decode(&entity)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if entity == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.UpdateUser(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update user"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
