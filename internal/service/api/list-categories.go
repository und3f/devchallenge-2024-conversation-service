package api

import (
	"encoding/json"
	"net/http"

	"devchallenge.it/conversation/internal/model"
)

func (s *Controller) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories := []model.Category{}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(categories)
}
