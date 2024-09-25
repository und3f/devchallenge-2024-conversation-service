package audio

import (
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
)

type Audio struct {
}

var SupportedAudioTypes = []string{"wav", "x-wav", "mpeg"}

func (audio *Audio) Download(audioUrl string) (bytes []byte, err error) {
	resp, err := http.Get(audioUrl)
	if err != nil {
		return bytes, err
	}

	ctHeader := resp.Header.Get("Content-Type")
	contentType := strings.Split(ctHeader, "/")
	if contentType[0] != "audio" {
		return nil, fmt.Errorf("Received wrong content type, audio type required: %s", ctHeader)
	}

	if !slices.Contains(SupportedAudioTypes, contentType[1]) {
		return nil, fmt.Errorf("Only wav and mp3 files are supported, content type suggested wrong type: %s", ctHeader)
	}

	defer resp.Body.Close()
	bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
