package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"golang-training-Theater/pkg/data"
)

type hallAPI struct {
	data *data.HallData
}

func (h hallAPI) createHall(writer http.ResponseWriter, request *http.Request) {
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
	id, err := h.data.AddHall(*entity)
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

func (h hallAPI) deleteHallById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Hall)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.data.DeleteHall(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete hall"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (h hallAPI) updateHall(writer http.ResponseWriter, request *http.Request) {
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
	err = h.data.UpdateHall(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update hall"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (h hallAPI) getHallById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Hall)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := h.data.FindByIdHall(*entity)
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
