package call

import (
	"devchallenge.it/conversation/internal/config"
	"devchallenge.it/conversation/internal/model"
	"devchallenge.it/conversation/internal/services"
	"devchallenge.it/conversation/internal/services/audio"
	"devchallenge.it/conversation/internal/services/nlp"
	"devchallenge.it/conversation/internal/services/whisper"
	"github.com/gorilla/mux"
)

type Controller struct {
	dao         *model.Dao
	analyzeChan chan AnalyzeTask

	nlp     *nlp.NLP
	whisper *whisper.Whisper
	audio   *audio.Audio
}

func Mount(r *mux.Router, dao *model.Dao, srvConf model.ServicesConf) {
	audio, whisper, nlp := services.CreateServices(srvConf)
	c := &Controller{
		dao:         dao,
		analyzeChan: make(chan AnalyzeTask, config.CALL_CHAN_SIZE),

		nlp:     nlp,
		audio:   audio,
		whisper: whisper,
	}

	go c.Analyzer()

	c.Mount(r)
}
