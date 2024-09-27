package call

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type CallResponse struct {
	Id int64 `json:"id"`

	Processed    bool    `json:"-"`
	ProcessError *string `json:"-"`

	Text          *string  `json:"text"`
	Name          *string  `json:"name"`
	Location      *string  `json:"location"`
	EmotionalTone *string  `json:"emotional_tone"`
	Categories    []string `json:"categories"`
}

func (c *Controller) GetCall(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	call, err := c.dao.GetCall(int64(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !call.Processed {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	if call.ProcessError != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{Error: *call.ProcessError})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(call)
}
