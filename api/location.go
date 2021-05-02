package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Alexandr59/golang-training-Theater/pkg/data"
)

type locationAPI struct {
	data *data.LocationData
}

func (l locationAPI) createLocation(writer http.ResponseWriter, request *http.Request) {
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
	id, err := l.data.AddLocation(*entity)
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

func (l locationAPI) deleteLocationById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Location)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := l.data.DeleteLocation(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete location"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (l locationAPI) updateLocation(writer http.ResponseWriter, request *http.Request) {
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
	err = l.data.UpdateLocation(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update location"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (l locationAPI) getLocationById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Location)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := l.data.FindByIdLocation(*entity)
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
