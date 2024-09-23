package api

import (
	"encoding/json"
	"log"
	"net/http"

	"devchallenge.it/conversation/internal/model"
)

func (s *Controller) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category

	log.Println(category)

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
