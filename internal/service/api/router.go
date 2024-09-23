package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) Mount(rootRouter *mux.Router) *mux.Router {
	r := rootRouter.PathPrefix(PATH_PREFIX).Subrouter()

	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		c.ListCategories(w, r)
	}).Methods("GET")

	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		c.CreateCategory(w, r)
	}).Methods("POST")

	return r
}
