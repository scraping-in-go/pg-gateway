package main

import (
	"github.com/deanishe/go-env"
	"github.com/gorilla/mux"
	"github.com/just1689/pg-gateway/db"
	"github.com/just1689/pg-gateway/web"
	"github.com/sirupsen/logrus"
	_ "go.uber.org/automaxprocs"
	"net/http"
	"os"
)

var poolSize = env.GetInt("poolSize")

func main() {

	checkEnvironmentVars()
	logrus.Println("Starting DB pool of size", poolSize)
	db.NextPoolCon = db.StartPool(poolSize)

	router := mux.NewRouter()
	router.HandleFunc("/", web.HandleOptions).Methods(http.MethodOptions)
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	router.HandleFunc("/{entity}", web.HandleGetAll).Methods(http.MethodGet)
	router.HandleFunc("/{entity}/{id}", web.HandleGet).Methods(http.MethodGet)
	router.HandleFunc("/{entity}/{field}/{id}", web.HandlePatch).Methods(http.MethodPatch)
	router.HandleFunc("/{entity}/{field}/{id}", web.HandleDelete).Methods(http.MethodDelete)
	router.HandleFunc("/{entity}/{field}/{id}", web.HandleGetMany).Methods(http.MethodGet)
	router.HandleFunc("/{entity}", web.HandleInsert).Methods(http.MethodPost)
	logrus.Println("Listening on", os.Getenv("listenAddr"))
	panic(http.ListenAndServe(os.Getenv("listenAddr"), router))
}

func checkEnvironmentVars() {
	if poolSize == 0 {
		logrus.Println("Setting poolSize to 1")
		poolSize = 1
	}

	//TODO: check and panic if unable to proceed
}
