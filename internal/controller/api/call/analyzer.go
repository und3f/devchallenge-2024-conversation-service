package call

import (
	"fmt"
	"log"

	"devchallenge.it/conversation/internal/model"
)

type AnalyzeTask struct {
	CallId int64
	Url    string
}

func (c *Controller) Analyzer() {

	for {
		select {
		case task := <-c.analyzeChan:
			c.ProcessCall(task.CallId, task.Url)
		}
	}
}

func (c *Controller) ProcessCall(callId int64, audioUrl string) {
	log.Printf("Processing call %d...", callId)

	call := c.AnalyzeCall(callId, audioUrl)

	c.dao.UpdateCall(call)

	err := "<nil>"
	if call.ProcessError != nil {
		err = *call.ProcessError
	}
	log.Printf("Processed call %d, err = %s", callId, err)
}

func (c *Controller) AnalyzeCall(callId int64, audioUrl string) model.Call {
	call := model.Call{Id: callId, Processed: true}

	audio, err := c.audio.Download(audioUrl)
	if err != nil {
		errStr := fmt.Sprintf("Failed to get audio file: %s", err)
		call.ProcessError = &errStr
		return call
	}

	log.Printf("  processing call %d: speech recognition...", callId)
	text, err := c.whisper.RecognizeSpeech(audio)
	if err != nil {
		errStr := fmt.Sprintf("Speech recongnition failure: %s", err)
		call.ProcessError = &errStr
		return call
	}

	call.Text = &text

	log.Printf("  processing call %d: sentiment prediction...", callId)
	emotional, err := c.nlp.GetSentiment(text)
	if err != nil {
		errStr := fmt.Sprintf("Failed to analyze emotional tone: %s", err)
		call.ProcessError = &errStr
		return call
	}

	call.EmotionalTone = &emotional

	log.Printf("  processing call %d: data extraction...", callId)
	data, err := c.nlp.ExtractData(text)
	if err != nil {
		errStr := fmt.Sprintf("Failed to extract data: %s", err)
		call.ProcessError = &errStr
		return call
	}

	if len(data.Name) > 0 {
		call.Name = &data.Name
	}

	if len(data.Name) > 0 {
		call.Location = &data.Location
	}

	return call
}
