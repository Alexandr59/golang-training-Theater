package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"golang-training-Theater/pkg/data"
)

type performanceAPI struct {
	data *data.PerformanceData
}

func (p performanceAPI) createPerformance(writer http.ResponseWriter, request *http.Request) {
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
	id, err := p.data.AddPerformance(*entity)
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

func (p performanceAPI) deletePerformanceById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Performance)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := p.data.DeletePerformance(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete performance"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (p performanceAPI) updatePerformance(writer http.ResponseWriter, request *http.Request) {
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
	err = p.data.UpdatePerformance(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update performance"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (p performanceAPI) getPerformanceById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Performance)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := p.data.FindByIdPerformance(*entity)
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
