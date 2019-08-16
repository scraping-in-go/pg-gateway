package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/scraping-in-go/svc-db-gateway/db"
	"github.com/scraping-in-go/svc-db-gateway/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	logrus.Println("Listening on", os.Getenv("listenAddr"))
	router := mux.NewRouter()

	router.HandleFunc("/{entity}", handleGetAll).Methods("GET")
	router.HandleFunc("/{entity}/{id}", handleGet).Methods("GET")
	router.HandleFunc("/{entity}", handleInsert).Methods("POST")
	panic(http.ListenAndServe(os.Getenv("listenAddr"), router))
}

func handleGetAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entity := vars["entity"]
	if entity == "" {
		http.Error(w, "You need to supply an entity: /{entity}", http.StatusBadRequest)
		return
	}

	c, err := db.GetEntityAll(entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "Application/json")
	rows := 0
	w.Write([]byte("["))
	for row := range c {
		rows++
		if rows > 1 {
			w.Write([]byte(","))
		}
		w.Write([]byte(row))

	}
	w.Write([]byte("]"))

}

func handleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "You need to supply an id: /{entity}/{id}", http.StatusBadRequest)
		return
	}

	entity := vars["entity"]
	if entity == "" {
		http.Error(w, "You need to supply an entity: /{entity}/{id}", http.StatusBadRequest)
		return
	}

	jsonS, err := db.GetEntityByID(entity, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "Application/json")
	w.Write([]byte(jsonS))

}

func handleInsert(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Could not read post body", http.StatusBadRequest)
		return
	}

	item := model.Insertable{}
	err = json.Unmarshal(b, &item)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Could not unmarshal item from body", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	entity := vars["entity"]
	if entity == "" {
		http.Error(w, "You need to supply an entity: /{entity}", http.StatusBadRequest)
		return
	}

	err = db.Insert(entity, item)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Could not insert", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Context-type", "Application/json")
	w.Write([]byte(`{"ok": true}`))

}
