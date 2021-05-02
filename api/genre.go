package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Alexandr59/golang-training-Theater/pkg/data"
)

type genreAPI struct {
	data *data.GenreData
}

func (g genreAPI) createGenre(writer http.ResponseWriter, request *http.Request) {
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
	id, err := g.data.AddGenre(*entity)
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

func (g genreAPI) deleteGenreById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Genre)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := g.data.DeleteGenre(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete genre"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (g genreAPI) updateGenre(writer http.ResponseWriter, request *http.Request) {
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
	err = g.data.UpdateGenre(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update genre"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (g genreAPI) getGenreById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Genre)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := g.data.FindByIdGenre(*entity)
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
