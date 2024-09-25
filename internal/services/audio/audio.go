package audio

import (
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

const MAX_LENGTH = 2726297

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

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return nil, fmt.Errorf("Invalid Content-Length value: %s", err)
	}
	if size > MAX_LENGTH {
		return nil, fmt.Errorf("File is too large: max allowed size %.1fMB", MAX_LENGTH/1024./1024.)
	}

	defer resp.Body.Close()
	bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
