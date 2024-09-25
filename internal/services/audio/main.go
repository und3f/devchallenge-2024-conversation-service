package audio

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

type Audio struct {
}

func (audio *Audio) Download(audioUrl string) (bytes []byte, err error) {
	resp, err := http.Get(audioUrl)
	if err != nil {
		return bytes, err
	}

	contentType := strings.Split(resp.Header.Get("Content-Type"), "/")
	if contentType[0] != "audio" {
		return nil, errors.New("Not audio file provided")
	}

	if !strings.Contains(contentType[1], "wav") {
		return nil, errors.New("Not wav audio url")
	}

	defer resp.Body.Close()
	bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
