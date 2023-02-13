package repository

import (
	"context"

	"github.com/muhangga/internal/entity"
)

func (r categoryRepository) Save(c entity.Category) (entity.Category, error) {
	sqlStmt := `INSERT INTO categories (name, icon) VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	_, err := r.db.ExecContext(ctx, sqlStmt, c.Name, c.Icon)

	if err != nil {
		return c, err
	}

	sqlStmt = `SELECT id, name, icon FROM categories ORDER BY id DESC LIMIT 1`

	row := r.db.QueryRowContext(ctx, sqlStmt)
	err = row.Scan(
		&c.ID,
		&c.Name,
		&c.Icon,
	)

	if err != nil {
		return c, err
	}

	return c, nil
}

func (r categoryRepository) FindAll() ([]entity.Category, error) {
	sqlStmt := `SELECT id, name, icon FROM categories`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, sqlStmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var c entity.Category
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.Icon,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, c)
	}
	return categories, nil

}

func (r categoryRepository) Update(c entity.Category) (entity.Category, error) {
	sqlStmt := `UPDATE categories SET name = $1, icon = $2, updated_at = $3 WHERE id = $4`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	_, err := r.db.ExecContext(ctx, sqlStmt, c.Name, c.Icon, c.UpdatedAt, c.ID)

	if err != nil {
		return c, err
	}

	sqlStmt = `SELECT id, name, icon FROM categories WHERE id = $1`

	row := r.db.QueryRowContext(ctx, sqlStmt, c.ID)
	err = row.Scan(
		&c.ID,
		&c.Name,
		&c.Icon,
	)

	if err != nil {
		return c, err
	}

	return c, nil
}

func (r categoryRepository) Delete(id int64) error {
	sqlStmt := `DELETE FROM categories WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	_, err := r.db.ExecContext(ctx, sqlStmt, id)

	if err != nil {
		return err
	}
	return nil
}
