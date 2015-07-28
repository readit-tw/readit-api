package main

import (
	"encoding/json"
	"github.com/readit-tw/readit-api/model"
	"github.com/readit-tw/readit-api/repository"
	"gopkg.in/mgo.v2"
	"net/http"
)

func createResourceHandler(rr repository.ResourceRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resource := &model.Resource{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&resource)

		if err != nil {
			http.Error(w, "Bad JSON Request", http.StatusBadRequest)
			return
		}

		created, err := rr.Create(resource)

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		resourceJson, err := json.Marshal(created)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(resourceJson)
	}
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("readit")
	resourceRepository := repository.NewMongoResourceRepository(db)

	http.HandleFunc("/resources", createResourceHandler(resourceRepository))
	http.ListenAndServe(":8080", nil)
}
