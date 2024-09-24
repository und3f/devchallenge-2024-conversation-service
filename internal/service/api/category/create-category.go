package category

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"devchallenge.it/conversation/internal/model"
)

const MIN_TITLE_LEN = 3

func (c *Controller) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if len(category.Title) < MIN_TITLE_LEN {
		err = errors.New(fmt.Sprintf("Title minimal length %d.", MIN_TITLE_LEN))
		log.Print(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	category, err = c.dao.CreateCategory(category)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
