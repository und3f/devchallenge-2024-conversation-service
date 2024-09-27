package whisper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

type WhisperResponse struct {
	Text  string  `json:"text"`
	Error *string `json:"error"`
}

func (whisper *Whisper) RecognizeSpeech(audio []byte) (text string, err error) {
	client := &http.Client{}

	url, err := url.JoinPath(whisper.Url, "/inference")
	if err != nil {
		return text, err
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	if fw, err := w.CreateFormFile("file", "file"); err != nil {
		return text, err
	} else {
		if _, err := io.Copy(fw, bytes.NewReader(audio)); err != nil {
			return text, err
		}
	}

	if fw, err := w.CreateFormField("response-format"); err != nil {
		return text, err
	} else {
		if _, err := io.Copy(fw, strings.NewReader("json")); err != nil {
			return text, err
		}
	}
	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return text, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		return text, err
	}

	if resp.StatusCode != http.StatusOK {
		return text, fmt.Errorf("Invalid response status: %s", resp.Status)
	}

	var whisperResp WhisperResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&whisperResp); err != nil {
		return text, err
	}

	if whisperResp.Error != nil {
		return text, fmt.Errorf("Analyzing error: %s", *whisperResp.Error)
	}

	return strings.TrimSpace(whisperResp.Text), nil
}
