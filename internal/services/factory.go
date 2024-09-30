package services

import (
	"devchallenge.it/conversation/internal/model"
	"devchallenge.it/conversation/internal/services/audio"
	"devchallenge.it/conversation/internal/services/nlp"
	"devchallenge.it/conversation/internal/services/whisper"
)

type ServicesFacade struct {
	Audio   *audio.Audio
	Whisper *whisper.Whisper
	NLP     *nlp.NLP
}

func CreateServices(srvConf model.ServicesConf) (*audio.Audio, *whisper.Whisper, *nlp.NLP) {
	return &audio.Audio{},
		&whisper.Whisper{srvConf.WhisperUrl},
		&nlp.NLP{Url: srvConf.NlpUrl}
}

func CreateServicesFacade(srvConf model.ServicesConf) ServicesFacade {
	audio, whisper, nlp := CreateServices(srvConf)
	return ServicesFacade{audio, whisper, nlp}
}
