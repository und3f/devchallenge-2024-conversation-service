package model

import "context"

type CallCreateResponse struct {
	Id int32 `json:"id"`
}

type Call struct {
	Id int32 `json:"id"`

	Processed    bool    `json:"-"`
	ProcessError *string `json:"-"`

	Text          *string  `json:"text"`
	Name          *string  `json:"name"`
	Location      *string  `json:"location"`
	EmotionalTone *string  `json:"emotional_tone"`
	Categories    []string `json:"categories,omitempty"`
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
		`
SELECT id, processed, error, name, location, emotional_tone, text
FROM calls WHERE id = $1
		`,
		id,
	).Scan(&call.Id, &call.Processed, &call.ProcessError, &call.Name, &call.Location, &call.EmotionalTone, &call.Text)
	if err != nil {
		return
	}

	if !call.Processed {
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
SELECT title
FROM categories
WHERE id IN (
	SELECT category_id
	FROM points
	JOIN category_points
		ON points.id = category_points.point_id
	WHERE
		LOWER((SELECT text FROM calls WHERE id = $1))
		LIKE CONCAT('%', CONCAT(LOWER(points.text), '%')))
ORDER BY title
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

func (d *Dao) UpdateCall(call Call) error {
	_, err := d.pg.Query(
		context.Background(),
		`
UPDATE calls
SET
	processed = $2,
	error = $3,
	text = $4,
	name = $5,
	location = $6,
	emotional_tone= $7
WHERE id = $1
	`,
		call.Id,
		call.Processed, call.ProcessError,
		call.Text, call.Name, call.Location, call.EmotionalTone,
	)

	return err
}
