package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Alexandr59/golang-training-Theater/pkg/data"
)

type userAPI struct {
	data *data.UserData
}

func (u userAPI) createUser(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.User)
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
	id, err := u.data.AddUser(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create user"))
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

func (u userAPI) deleteUserById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.User)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := u.data.DeleteUser(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete user"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (u userAPI) updateUser(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.User)
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
	err = u.data.UpdateUser(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update user"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (u userAPI) getUserById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.User)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := u.data.FindByIdUser(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get user"))
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

func (u userAPI) getAllUsers(writer http.ResponseWriter, request *http.Request) {
	account := new(data.Account)

	if n, err := strconv.Atoi(request.URL.Query().Get("idAccount")); err == nil {
		account.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := u.data.ReadAllUsers(*account)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get users"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
