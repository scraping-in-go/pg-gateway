package main

import (
	"github.com/deanishe/go-env"
	"github.com/gorilla/mux"
	"github.com/scraping-in-go/svc-db-gateway/db"
	"github.com/scraping-in-go/svc-db-gateway/web"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var poolSize = env.GetInt("poolSize")

func main() {
	checkEnvironmentVars()
	logrus.Println("Starting DB pool of size", poolSize)
	db.NextPoolCon = db.StartPool(poolSize)

	router := mux.NewRouter()
	router.HandleFunc("/{entity}", web.HandleGetAll).Methods("GET")
	router.HandleFunc("/{entity}/{id}", web.HandleGet).Methods("GET")
	router.HandleFunc("/{entity}/{field}/{id}", web.HandleGetMany).Methods("GET")
	router.HandleFunc("/{entity}", web.HandleInsert).Methods("POST")
	logrus.Println("Listening on", os.Getenv("listenAddr"))
	panic(http.ListenAndServe(os.Getenv("listenAddr"), router))
}

func checkEnvironmentVars() {
	if poolSize == 0 {
		logrus.Println("Setting poolSize to 1")
	}

	//TODO: check and panic if unable to proceed
}
