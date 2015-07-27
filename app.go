package main

import (
	"encoding/json"
	"github.com/readit-tw/readit-api/model"
	"net/http"
)

func createResourceHandler(w http.ResponseWriter, r *http.Request) {
	resource := &model.Resource{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&resource)
	if err != nil {
		http.Error(w, "Bad JSON Request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resourceJson, err := json.Marshal(resource)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(resourceJson)
}

func main() {
	http.HandleFunc("/resources", createResourceHandler)
	http.ListenAndServe(":8080", nil)
}
