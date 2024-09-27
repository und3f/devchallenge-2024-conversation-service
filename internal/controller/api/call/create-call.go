package call

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"devchallenge.it/conversation/internal/model"
	"github.com/go-http-utils/headers"
)

type CallCreateRequest struct {
	AudioUrl string `json:"audio_url"`
}

func (c *Controller) CreateCall(w http.ResponseWriter, r *http.Request) {
	var callCreate CallCreateRequest

	err := json.NewDecoder(r.Body).Decode(&callCreate)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	callId, err := c.dao.CreateCall(callCreate.AudioUrl)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	call := model.CallCreateResponse{
		Id: callId,
	}

	if err := c.scheduleCallAnalyze(callId, callCreate.AudioUrl); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusTooManyRequests)
		return
	}

	w.Header().Set(headers.Location, "/api/call/"+strconv.Itoa(int(callId)))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(call)
}

func (c *Controller) scheduleCallAnalyze(callId int64, url string) error {
	select {
	case c.analyzeChan <- AnalyzeTask{CallId: callId, Url: url}:
	default:
		return errors.New("Failed to schedule task as waiting queue is full.")
	}

	return nil
}
