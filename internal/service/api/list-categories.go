package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (c *Controller) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.dao.ListCategories()
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(categories)
}
