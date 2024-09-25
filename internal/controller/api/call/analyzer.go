package call

import (
	"fmt"
	"log"

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
	call := model.Call{Id: callId, Processed: true}

	audio, err := c.audio.Download(audioUrl)
	if err != nil {
		errStr := fmt.Sprintf("Failed to get audio file: %s", err)
		call.ProcessError = &errStr
		return call
	}

	text, err := c.whisper.RecognizeSpeech(callId, audio)
	if err != nil {
		errStr := fmt.Sprintf("Speech recongnition failure: %s", err)
		call.ProcessError = &errStr
		return call
	}

	call.Text = &text

	emotional, err := c.nlp.GetSentiment(text)
	if err != nil {
		errStr := fmt.Sprintf("Failed to analyze emotional tone: %s", err)
		call.ProcessError = &errStr
		return call
	}

	call.EmotionalTone = &emotional

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