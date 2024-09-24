package call

import (
	"devchallenge.it/conversation/internal/model"
	"github.com/gorilla/mux"
)

func Mount(r *mux.Router, dao *model.Dao) {
	c := &Controller{dao: dao}
	c.Mount(r)
}

type Controller struct {
	dao            *model.Dao
	subscribeRoute *mux.Route
}
