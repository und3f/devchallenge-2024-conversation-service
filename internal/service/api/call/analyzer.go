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

type SentimentTextPayload struct {
	Text   string            `json:"text"`
	Probas map[string]string `json:"probas"`
}

type SentimentResp struct {
	Output string `json:"output"`
}

var EmotionOutputToEmotionValue = map[string]string{
	"NEG": "Negative",
	"POS": "Positive",
	"NEU": "Neutral",
}

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

	emotional, err := c.RequestSentimentAnalyze(text)
	if err != nil {
		errStr := fmt.Sprintf("Failed to analyze emotional tone: %s", err)
		call.ProcessError = &errStr
		return call
	}

	call.EmotionalTone = &emotional

	return call
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

	url, err := url.JoinPath(c.srvConf.WhisperUrl, "/inference")
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

func (c *Controller) RequestSentimentAnalyze(text string) (sentiment string, err error) {
	url, err := url.JoinPath(c.srvConf.SentimentUrl, "/emotion")
	if err != nil {
		return sentiment, err
	}

	request := SentimentTextPayload{Text: text}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(&request)

	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		return sentiment, err
	}

	if resp.StatusCode != http.StatusOK {
		return sentiment, fmt.Errorf("Invalid response status: %s", resp.Status)
	}

	var sentimentResp SentimentResp
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&sentimentResp); err != nil {
		return sentiment, err
	}

	emotion, ok := EmotionOutputToEmotionValue[sentimentResp.Output]
	if !ok {
		return sentiment, fmt.Errorf("Unexpected emotion value: %s", sentimentResp.Output)
	}

	return emotion, nil
}
