package model

import (
	"context"
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/jackc/pgx/v5"
)

var ErrDupCategoryTitle = errors.New("Duplicate category title.")

type Category struct {
	Id     int64    `json:"id"`
	Title  string   `json:"title"`
	Points []string `json:"points,omitempty"`
}

func (d *Dao) ListCategories() ([]Category, error) {
	categories := make([]Category, 0)

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

		id := values[0].(int64)
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

func (d *Dao) GetCategoryPoints(id int64) ([]string, error) {
	points := []string{}

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

	slices.SortFunc(points, strings.Compare)

	return points, nil
}

func (d *Dao) CreateCategory(createReq Category) (category Category, err error) {
	var totalFoundCategories int64
	err = d.pg.QueryRow(
		context.Background(),
		"SELECT COUNT(*) FROM categories WHERE title = $1",
		createReq.Title,
	).Scan(&totalFoundCategories)
	if totalFoundCategories > 0 {
		return category, fmt.Errorf("Category \"%s\" already exists.",
			createReq.Title)
	}

	tx, err := d.pg.BeginTx(
		context.Background(),
		pgx.TxOptions{},
	)
	if err != nil {
		return category, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		} else {
			tx.Commit(context.Background())
		}
	}()

	var id int64
	err = tx.QueryRow(
		context.Background(),
		"INSERT INTO categories (title) VALUES($1) RETURNING id",
		createReq.Title,
	).Scan(&id)

	if err != nil {
		return
	}

	if err = d.BindCategoryPoints(tx, id, createReq.Points); err != nil {
		return
	}

	category = createReq
	category.Id = id

	return category, nil
}

func (d *Dao) BindCategoryPoints(tx pgx.Tx, category_id int64, points []string) error {
	for _, point := range points {
		pointId, err := d.CreateOrGetPoint(point)
		if err != nil {
			return err
		}

		if err := d.AddCategoryPoint(tx, category_id, pointId); err != nil {
			return err
		}
	}

	return nil
}

func (d *Dao) CreateOrGetPoint(text string) (id int64, err error) {
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

func (d *Dao) AddCategoryPoint(tx pgx.Tx, categoryId int64, pointId int64) (err error) {
	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO category_points (category_id, point_id) VALUES ($1, $2)",
		categoryId, pointId,
	)
	return
}

func (d *Dao) UpdateCategory(newCategoryValue Category) (category *Category, err error) {
	tx, err := d.pg.BeginTx(
		context.Background(),
		pgx.TxOptions{},
	)
	if err != nil {
		return category, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		} else {
			tx.Commit(context.Background())
		}
	}()

	if len(newCategoryValue.Title) > 0 {
		cmd, err := tx.Exec(
			context.Background(),
			"UPDATE categories SET title = $2 WHERE id = $1",
			newCategoryValue.Id,
			newCategoryValue.Title,
		)
		if err != nil {
			return nil, err
		}

		if cmd.RowsAffected() == 0 {
			log.Printf("UpdateCategory %d category not found.", category.Id)
			return nil, nil
		}
	} else {
		err := d.pg.QueryRow(
			context.Background(),
			"SELECT title FROM categories WHERE id = $1",
			newCategoryValue.Id,
		).Scan(&newCategoryValue.Title)

		if err != nil {
			return nil, err
		}
	}

	if len(newCategoryValue.Points) > 0 {
		_, err = tx.Exec(
			context.Background(),
			"DELETE FROM category_points WHERE category_id = $1",
			newCategoryValue.Id,
		)
		if err != nil {
			return
		}

		err = d.BindCategoryPoints(tx, newCategoryValue.Id, newCategoryValue.Points)
		if err != nil {
			return
		}
	} else {
		newCategoryValue.Points = make([]string, 0)
	}

	slices.SortFunc(newCategoryValue.Points, strings.Compare)
	return &newCategoryValue, nil
}

func (d *Dao) DeleteCategory(categoryId int64) (deleted bool, err error) {
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
