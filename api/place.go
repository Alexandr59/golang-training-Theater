package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"golang-training-Theater/pkg/data"
)

type placeAPI struct {
	data *data.PlaceData
}

func (p placeAPI) createPlace(writer http.ResponseWriter, request *http.Request) {
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
	id, err := p.data.AddPlace(*entity)
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

func (p placeAPI) deletePlaceById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Place)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := p.data.DeletePlace(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete place"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (p placeAPI) updatePlace(writer http.ResponseWriter, request *http.Request) {
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
	err = p.data.UpdatePlace(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update place"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (p placeAPI) getPlaceById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Place)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := p.data.FindByIdPlace(*entity)
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
