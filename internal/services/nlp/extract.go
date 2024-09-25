package nlp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type DataExtractionRequest struct {
	Text   string `json:"text"`
	Schema string `json:"schema"`
}

type DataExtractionSchema struct {
	Location string `json:"location"`
	Name     string `json:"name"`
}

func (nlp *NLP) ExtractData(text string) (data DataExtractionSchema, err error) {
	url, err := url.JoinPath(nlp.Url, "/extract")
	if err != nil {
		return data, err
	}

	var schema DataExtractionSchema

	schemaBytes, err := json.Marshal(&schema)
	if err != nil {
		return data, err
	}

	request := DataExtractionRequest{
		Text:   text,
		Schema: string(schemaBytes),
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(&request)

	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		return data, err
	}

	if resp.StatusCode != http.StatusOK {
		return data, fmt.Errorf("Invalid response status: %s", resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}
