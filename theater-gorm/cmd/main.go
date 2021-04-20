package main

import (
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"os"

	"golang-training-Theater/theater-gorm/api"
	//"github.com/Alexandr59/golang-training-Theater/theater-gorm/api"
	"github.com/Alexandr59/golang-training-Theater/theater-gorm/pkg/data"
	"github.com/Alexandr59/golang-training-Theater/theater-gorm/pkg/db"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "Theater_db"
	}
	if password == "" {
		password = "5959"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}

func main() {
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}
	theaterData := data.NewTheaterData(conn)

	r := mux.NewRouter()
	api.ServerTheaterResource(r, *theaterData)
	r.Use(mux.CORSMethodMiddleware(r))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Server doesn't listen port...", err)
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...", err)
	}

	//err = theaterData.UpdateAccount(data.Account{
	//	Id:          11,
	//	FirstName:   "YYYYYYY",
	//	LastName:    "OOOOOOOOO",
	//	PhoneNumber: "435345435e345",
	//	Email:       " lvlf,d;d",
	//})
	//if err != nil {
	//	log.Fatalf("got an error when tried to call UpdateAccount method: %v", err)
	//}
	//
	//err = theaterData.DeleteAccount(data.Account{Id: 13})
	//if err != nil {
	//	log.Fatalf("got an error when tried to call DeleteAccount method: %v", err)
	//}
	//
	//_, err = theaterData.AddAccount(data.Account{
	//	FirstName:   "Dim",
	//	LastName:    "Ivanov",
	//	PhoneNumber: "+375296574897",
	//	Email:       "dimaivanov@gmail.com",
	//})
	//if err != nil {
	//	log.Fatalf("got an error when tried to call AddAccount method: %v", err)
	//}
	//a, err := theaterData.FindByIdAccount(data.Account{Id: 1})
	//
	//fmt.Println(a)

	//tickets, err := theaterData.ReadAllTickets()
	//if err != nil {
	//	log.Fatalf("got an error when tried to call ReadAllTickets method: %v", err)
	//}
	//for _, el := range tickets {
	//	fmt.Println(el)
	//}
	//
	//posters, err := theaterData.ReadAllPosters()
	//if err != nil {
	//	log.Fatalf("got an error when tried to call ReadAllPosters method: %v", err)
	//}
	//for _, el := range posters {
	//	log.Println(el)
	//}
	//
	//users, err := theaterData.ReadAllUsers(data.Account{Id: 1})
	//if err != nil {
	//	log.Fatalf("got an error when tried to call ReadAllUsers method: %v", err)
	//}
	//for _, el := range users {
	//	log.Println(el)
	//}
}
