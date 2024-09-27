package category

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"devchallenge.it/conversation/internal/model"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
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

	if len(category.Points) < 1 {
		err = errors.New(fmt.Sprintf("At least one category point should be provided"))
		log.Print(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	category, err = c.dao.CreateCategory(category)
	if err != nil {
		status := http.StatusInternalServerError

		if isConstraintError(err) {
			status = http.StatusUnprocessableEntity
		}

		log.Print(err)
		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func isConstraintError(err error) bool {
	var pgErr *pgconn.PgError
	return (errors.As(err, &pgErr)) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code)
}
