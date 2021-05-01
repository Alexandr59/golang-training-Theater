package main

import (
	api2 "golang-training-Theater/api"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	//"github.com/Alexandr59/golang-training-Theater/api"
	"github.com/Alexandr59/golang-training-Theater/pkg/data"
	"github.com/Alexandr59/golang-training-Theater/pkg/db"
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
	api2.ServerTheaterResource(r, *theaterData)
	r.Use(mux.CORSMethodMiddleware(r))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Server doesn't listen port...", err)
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...", err)
	}
}
