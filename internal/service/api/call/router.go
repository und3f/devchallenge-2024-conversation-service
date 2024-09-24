package call

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PATH_PREFIX = "/call"
)

func (c *Controller) Mount(rootRouter *mux.Router) *mux.Router {
	r := rootRouter.PathPrefix(PATH_PREFIX).Subrouter()

	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		c.CreateCall(w, r)
	}).Methods("POST")

	r.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		c.GetCall(w, r)
	}).Methods("GET")

	return r
}
