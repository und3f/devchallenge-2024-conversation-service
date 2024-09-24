package model

import "context"

type CallCreateResponse struct {
	Id int32 `json:"id"`
}

type Call struct {
	Id            int32      `json:"id"`
	Processed     bool       `json:"-"`
	Name          *string    `json:"name"`
	Location      *string    `json:"location"`
	EmotionalTone *string    `json:"emotional_tone"`
	Text          *string    `json:"text"`
	Categories    []Category `json:"categories"`
}

func (d *Dao) CreateCall(audioFile string) (id int32, err error) {
	err = d.pg.QueryRow(
		context.Background(),
		"INSERT INTO calls DEFAULT VALUES RETURNING id",
	).Scan(&id)

	if err != nil {
		return
	}

	return
}

func (d *Dao) GetCall(id int32) (call Call, err error) {
	err = d.pg.QueryRow(
		context.Background(),
		"SELECT * FROM calls WHERE id = $1",
		id,
	).Scan(&call.Id, &call.Processed, &call.Name, &call.Location, &call.EmotionalTone, &call.Text)

	return
}
