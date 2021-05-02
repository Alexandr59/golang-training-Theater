package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Alexandr59/golang-training-Theater/pkg/data"
)

type posterAPI struct {
	data *data.PosterData
}

func (p posterAPI) createPoster(writer http.ResponseWriter, request *http.Request) {
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
	id, err := p.data.AddPoster(*entity)
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

func (p posterAPI) deletePosterById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Poster)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := p.data.DeletePoster(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete poster"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (p posterAPI) updatePoster(writer http.ResponseWriter, request *http.Request) {
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
	err = p.data.UpdatePoster(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update poster"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (p posterAPI) getPosterById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Poster)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := p.data.FindByIdPoster(*entity)
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

func (p posterAPI) getAllPosters(writer http.ResponseWriter, _ *http.Request) {
	posters, err := p.data.ReadAllPosters()
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
