package call

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"devchallenge.it/conversation/internal/model"
)

func (c *Controller) ProcessCall(callId int32, audioUrl string) {
	c.analyzeMutex.Lock()
	defer c.analyzeMutex.Unlock()

	log.Printf("Processing call %d...", callId)

	call := c.AnalyzeCall(callId, audioUrl)

	c.dao.UpdateCall(call)

	err := "<nil>"
	if call.ProcessError != nil {
		err = *call.ProcessError
	}
	log.Printf("Processed call %d, err = %s", callId, err)
}

func (c *Controller) AnalyzeCall(callId int32, audioUrl string) model.Call {
	name := strings.Split(audioUrl[strings.LastIndex(audioUrl, "/")+1:], ".")[0]

	call := model.Call{Id: callId, Processed: true, Name: &name}

	audio, err := c.GetAudioFile(audioUrl)
	if err != nil {
		errStr := fmt.Sprintf("Failed to get audio file: %s", err)
		call.ProcessError = &errStr
		return call
	}

	text, err := c.RequestAudioAnalyze(callId, audio)
	if err != nil {
		errStr := fmt.Sprintf("Failed to analyze audio: %s", err)
		call.ProcessError = &errStr
		return call
	}

	call.Text = &text
	call.Categories, err = c.AnalyzeTextCategories(text)
	if err != nil {
		errStr := fmt.Sprintf("Failed to categorize audio: %s", err)
		call.ProcessError = &errStr
		return call
	}

	return call
}

func (c *Controller) AnalyzeTextCategories(text string) (categories []string, err error) {
	return nil, nil
}

func (c *Controller) GetAudioFile(audioUrl string) (audio []byte, err error) {
	resp, err := http.Get(audioUrl)
	if err != nil {
		return audio, err
	}

	contentType := strings.Split(resp.Header.Get("Content-Type"), "/")
	if contentType[0] != "audio" {
		return nil, errors.New("Not audio file provided")
	}

	if !strings.Contains(contentType[1], "wav") {
		return nil, errors.New("Not wav audio url")
	}

	defer resp.Body.Close()
	audio, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func (c *Controller) RequestAudioAnalyze(callId int32, audio []byte) (text string, err error) {
	client := &http.Client{}

	url, err := url.JoinPath(c.whisperUrl, "/inference")
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
