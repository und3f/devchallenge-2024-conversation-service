package call

import (
	"devchallenge.it/conversation/internal/config"
	"devchallenge.it/conversation/internal/model"
	"devchallenge.it/conversation/internal/services"
	"github.com/gorilla/mux"
)

type Controller struct {
	dao         *model.Dao
	analyzeChan chan AnalyzeTask

	srv  services.ServicesFacade
	quit chan any
}

func Mount(r *mux.Router, dao *model.Dao, srv services.ServicesFacade, quit chan any) {
	c := &Controller{
		dao:         dao,
		analyzeChan: make(chan AnalyzeTask, config.CALL_CHAN_SIZE),

		srv:  srv,
		quit: quit,
	}

	go c.Analyzer()

	c.Mount(r)
}
