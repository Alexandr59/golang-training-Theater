package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alexandr59/golang-training-Theater/theater-gorm/pkg/data"
	"github.com/gorilla/mux"
)

type theaterAPI struct {
	data *data.TheaterData
}

func ServerTheaterResource(r *mux.Router, data data.TheaterData) {
	api := &theaterAPI{data: &data}
	r.HandleFunc("/tickets", api.getAllTickets).Methods("GET")
	//r.HandleFunc("/users", api.createUser).Methods("POST")
}

func (a theaterAPI) getAllTickets(writer http.ResponseWriter, request *http.Request) {
	users, err := a.data.ReadAllTickets()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get tickets"))
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

//func (a userAPI) createUser(writer http.ResponseWriter, request *http.Request) {
//	user := new(data.User)
//	err := json.NewDecoder(request.Body).Decode(&user)
//	if err != nil {
//		log.Printf("failed reading JSON: %s\n", err)
//		writer.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	if user == nil {
//		log.Printf("failed empty JSON\n")
//		writer.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	err = a.data.Add(*user)
//	if err != nil {
//		log.Println("user hasn't been created")
//		writer.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	writer.WriteHeader(http.StatusCreated)
//}
