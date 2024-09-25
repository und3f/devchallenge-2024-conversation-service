package category

import (
	"devchallenge.it/conversation/internal/model"
	"github.com/gorilla/mux"
)

type Controller struct {
	dao *model.Dao
}

func Mount(r *mux.Router, dao *model.Dao) {
	c := &Controller{dao: dao}
	c.Mount(r)
}
