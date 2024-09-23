package api

import (
	"fmt"
	"net/http"

	"devchallenge.it/conversation/internal/model"
	"github.com/gorilla/mux"
)

type Controller struct {
	dao            *model.Dao
	subscribeRoute *mux.Route
}

func New(r *mux.Router, dao *model.Dao) *Controller {
	s := &Controller{dao: dao}
	s.Mount(r)
	return s
}

func (s *Controller) Mount(r *mux.Router) *mux.Router {
	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		s.root(w, r)
	})

	return r
}

func (s *Controller) root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello there!")
}
