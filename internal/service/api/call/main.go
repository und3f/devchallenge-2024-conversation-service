package call

import (
	"devchallenge.it/conversation/internal/model"
	"github.com/gorilla/mux"
)

type Controller struct {
	dao        *model.Dao
	whisperUrl string
}

func Mount(r *mux.Router, dao *model.Dao, whisperUrl string) {
	c := &Controller{dao: dao, whisperUrl: whisperUrl}
	c.Mount(r)
}
