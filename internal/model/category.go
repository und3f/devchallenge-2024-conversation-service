package model

import (
	"context"
)

type Category struct {
	Id     int32    `json:"id",omitifempty`
	Title  string   `json:"title"`
	Points []string `json:"points",omitifempty`
}

func (d *Dao) ListCategories() ([]Category, error) {
	var categories []Category

	rows, err := d.pg.Query(
		context.Background(),
		`SELECT id, title FROM categories`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		id := values[0].(int32)
		title := values[1].(string)
		points, err := d.GetCategoryPoints(id)
		if err != nil {
			return nil, err
		}

		categories = append(categories, Category{
			Id:     id,
			Title:  title,
			Points: points,
		})
	}

	return categories, nil
}

func (d *Dao) GetCategoryPoints(id int32) ([]string, error) {
	var points []string

	rows, err := d.pg.Query(
		context.Background(),
		`
SELECT points.text
FROM category_points
JOIN
	points ON category_points.point_id = points.id
WHERE
	category_points.category_id = $1
		`,
		id,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		text := values[0].(string)
		points = append(points, text)
	}

	return points, nil
}

func (d *Dao) CreateCategory(createReq Category) (category Category, err error) {
	var id int32
	err = d.pg.QueryRow(
		context.Background(),
		"INSERT INTO categories (title) VALUES($1) RETURNING id",
		createReq.Title,
	).Scan(&id)

	if err != nil {
		return
	}

	for _, point := range createReq.Points {
		pointId, err := d.CreateOrGetPoint(point)

		if err != nil {
			return category, err
		}
		if err := d.AddCategoryPoint(id, pointId); err != nil {
			return category, err
		}
	}

	category = createReq
	category.Id = id

	return category, nil
}

func (d *Dao) CreateOrGetPoint(text string) (id int32, err error) {
	err = d.pg.QueryRow(
		context.Background(),
		"SELECT id FROM points WHERE text = $1",
		text,
	).Scan(&id)

	if err == nil {
		return id, err
	}

	err = d.pg.QueryRow(
		context.Background(), "INSERT INTO points (text) VALUES ($1) RETURNING id;", text).Scan(&id)

	if err != nil {
		return
	}

	return
}

func (d *Dao) AddCategoryPoint(categoryId int32, pointId int32) (err error) {
	_, err = d.pg.Exec(
		context.Background(),
		"INSERT INTO category_points (category_id, point_id) VALUES ($1, $2)",
		categoryId, pointId,
	)
	return
}

func (d *Dao) DeleteCategory(categoryId int32) (deleted bool, err error) {
	cmd, err := d.pg.Exec(
		context.Background(),
		"DELETE FROM categories WHERE id = $1",
		categoryId,
	)

	if err == nil {
		deleted = cmd.RowsAffected() > 0
	}
	return
}
