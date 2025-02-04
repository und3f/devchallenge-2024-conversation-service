package category

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"devchallenge.it/conversation/internal/model"
	"github.com/gorilla/mux"
)

func (c *Controller) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var newCategoryValue model.Category
	err = json.NewDecoder(r.Body).Decode(&newCategoryValue)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	newCategoryValue.Id = int64(id)

	category, err := c.dao.UpdateCategory(newCategoryValue)
	if err != nil {
		status := http.StatusInternalServerError
		if isConstraintError(err) {
			status = http.StatusUnprocessableEntity
		}

		log.Print(err)
		http.Error(w, err.Error(), status)
		return
	}

	if category == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}
