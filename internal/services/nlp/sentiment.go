package nlp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func (nlp *NLP) GetSentiment(text string) (sentiment string, err error) {
	url, err := url.JoinPath(nlp.Url, "/emotion")
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
