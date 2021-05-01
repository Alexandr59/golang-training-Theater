package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"golang-training-Theater/pkg/data"
)

type ticketAPI struct {
	data *data.TicketData
}

func (t ticketAPI) createTicket(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Ticket)
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
	id, err := t.data.AddTicket(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to create ticket"))
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

func (t ticketAPI) deleteTicketById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Ticket)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := t.data.DeleteTicket(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to delete ticket"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (t ticketAPI) updateTicket(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Ticket)
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
	err = t.data.UpdateTicket(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to update ticket"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (t ticketAPI) getTicketById(writer http.ResponseWriter, request *http.Request) {
	entity := new(data.Ticket)
	if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
		entity.Id = n
	} else {
		log.Printf("failed reading id: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	entry, err := t.data.FindByIdTicket(*entity)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get ticket"))
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

func (t ticketAPI) getAllTickets(writer http.ResponseWriter, _ *http.Request) {
	tickets, err := t.data.ReadAllTickets()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get tickets"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(tickets)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
