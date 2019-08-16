package main

import (
	"encoding/json"
	"github.com/google/uuid"
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
	router.HandleFunc("/insert", handleInsert)
	router.HandleFunc("/get/{id}", handleGet)
	panic(http.ListenAndServe(os.Getenv("listenAddr"), router))
}

func handleInsert(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Could not read post body", http.StatusBadRequest)
		return
	}
	entity := model.Entity{}
	err = json.Unmarshal(b, &entity)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Could not unmarshal entity from body", http.StatusBadRequest)
		return
	}

	if entity.ID == "" {
		entity.ID = uuid.New().String()
	}
	err = db.Insert(entity)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Could not insert", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Context-type", "Application/json")
	w.Write([]byte(`{"ok": true}`))

}

func handleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "You need to supply an id: /get/{id}", http.StatusBadRequest)
		return
	}

	entity, err := db.GetEntityByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(entity)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Context-type", "Application/json")
	w.Write(b)

}
