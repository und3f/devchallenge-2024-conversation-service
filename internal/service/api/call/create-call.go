package call

import (
	"encoding/json"
	"log"
	"net/http"

	"devchallenge.it/conversation/internal/model"
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(call)
}
