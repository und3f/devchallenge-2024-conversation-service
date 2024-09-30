package services

import (
	"devchallenge.it/conversation/internal/model"
	"devchallenge.it/conversation/internal/services/audio"
	"devchallenge.it/conversation/internal/services/nlp"
	"devchallenge.it/conversation/internal/services/whisper"
)

func CreateServices(srvConf model.ServicesConf) (*audio.Audio, *whisper.Whisper, *nlp.NLP) {
	return &audio.Audio{},
		&whisper.Whisper{srvConf.WhisperUrl},
		&nlp.NLP{Url: srvConf.NlpUrl}
}
