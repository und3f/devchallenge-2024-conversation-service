package call

import (
	"sync"

	"devchallenge.it/conversation/internal/model"
	"github.com/gorilla/mux"
)

type Controller struct {
	srvConf      model.ServicesConf
	dao          *model.Dao
	analyzeMutex sync.Mutex
}

func Mount(r *mux.Router, dao *model.Dao, srvConf model.ServicesConf) {
	c := &Controller{dao: dao, srvConf: srvConf}
	c.Mount(r)
}
