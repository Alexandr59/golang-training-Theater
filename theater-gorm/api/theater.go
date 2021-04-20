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
	//r.HandleFunc("/users", api.createUser).Methods("POST")
}

func (a theaterAPI) getAllTickets(writer http.ResponseWriter, _ *http.Request) {
	tickets, err := a.data.ReadAllTickets()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get tickets"))
		if err != nil {
			log.Println(err)
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
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := a.data.ReadAllUsers(*account)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get users"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdAccount(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get account"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdGenre(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get genre"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdHall(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get hall"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdLocation(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get location"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdPerformance(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get performance"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdPlace(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get place"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdPoster(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get poster"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdPrice(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get price"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdRole(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get role"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdSchedule(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get schedule"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdSector(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get sector"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdTicket(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get ticket"))
		if err != nil {
			log.Println(err)
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
		log.Printf("failed reading id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := a.data.FindByIdUser(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get user"))
		if err != nil {
			log.Println(err)
		}
	}
	err = json.NewEncoder(writer).Encode(entry)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
