package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Alexandr59/golang-training-Theater/pkg/data"
)

type scheduleAPI struct {
	data *data.ScheduleData
}

func (s scheduleAPI) createSchedule(writer http.ResponseWriter, request *http.Request) {
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
	id, err := s.data.AddSchedule(*entity)
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

func (s scheduleAPI) deleteScheduleById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Schedule)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := s.data.DeleteSchedule(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete schedule"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s scheduleAPI) updateSchedule(writer http.ResponseWriter, request *http.Request) {
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
	err = s.data.UpdateSchedule(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update schedule"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s scheduleAPI) getScheduleById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Schedule)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := s.data.FindByIdSchedule(*entity)
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
