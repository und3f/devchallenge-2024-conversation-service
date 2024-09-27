package call

import (
	"devchallenge.it/conversation/internal/model"
	"devchallenge.it/conversation/internal/services/audio"
	"devchallenge.it/conversation/internal/services/nlp"
	"devchallenge.it/conversation/internal/services/whisper"
	"github.com/gorilla/mux"
)

type Controller struct {
	dao         *model.Dao
	analyzeChan chan AnalyzeTask

	nlp     nlp.NLP
	whisper whisper.Whisper
	audio   audio.Audio
}

func Mount(r *mux.Router, dao *model.Dao, srvConf model.ServicesConf) {
	c := &Controller{
		dao:         dao,
		analyzeChan: make(chan AnalyzeTask, 10),

		nlp:     nlp.NLP{srvConf.NlpUrl},
		audio:   audio.Audio{},
		whisper: whisper.Whisper{srvConf.WhisperUrl},
	}

	go c.Analyzer()

	c.Mount(r)
}
