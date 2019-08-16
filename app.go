package main

import (
	"github.com/gorilla/mux"
	"github.com/scraping-in-go/svc-db-gateway/web"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var LB = []byte("[")
var RB = []byte("]")
var COMMA = []byte(",")

func main() {

	logrus.Println("Listening on", os.Getenv("listenAddr"))
	router := mux.NewRouter()

	router.HandleFunc("/{entity}", web.HandleGetAll).Methods("GET")
	router.HandleFunc("/{entity}/{id}", web.HandleGet).Methods("GET")
	router.HandleFunc("/{entity}/{field}/{id}", web.HandleGetMany).Methods("GET")
	router.HandleFunc("/{entity}", web.HandleInsert).Methods("POST")
	panic(http.ListenAndServe(os.Getenv("listenAddr"), router))
}
