package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"golang-training-Theater/pkg/data"
)

type sectorAPI struct {
	data *data.SectorData
}

func (s sectorAPI) createSector(writer http.ResponseWriter, request *http.Request) {
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
	id, err := s.data.AddSector(*entity)
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

func (s sectorAPI) deleteSectorById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Sector)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := s.data.DeleteSector(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete sector"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s sectorAPI) updateSector(writer http.ResponseWriter, request *http.Request) {
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
	err = s.data.UpdateSector(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update sector"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s sectorAPI) getSectorById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Sector)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := s.data.FindByIdSector(*entity)
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
