package model

import "context"

type CallCreateResponse struct {
	Id int32 `json:"id"`
}

type Call struct {
	Id            int32    `json:"id"`
	Processed     bool     `json:"-"`
	Name          *string  `json:"name"`
	Location      *string  `json:"location"`
	EmotionalTone *string  `json:"emotional_tone"`
	Text          *string  `json:"text"`
	Categories    []string `json:"categories"`
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
	if err != nil {
		return
	}

	call.Categories, err = d.GetCallCategories(id)
	if err != nil {
		return
	}

	return
}

func (d *Dao) GetCallCategories(callId int32) (categories []string, err error) {
	rows, err := d.pg.Query(
		context.Background(),
		`
SELECT categories.title
FROM call_categories
JOIN
	categories ON call_categories.category_id = categories.id
WHERE
	call_categories.call_id = $1
ORDER BY categories.title ASC
		`,
		callId,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		title := values[0].(string)
		categories = append(categories, title)
	}

	return
}
