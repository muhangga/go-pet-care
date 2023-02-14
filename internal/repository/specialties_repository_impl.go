package repository

import (
	"context"

	"github.com/muhangga/internal/entity"
)

func (r *specialtiesRepository) FindAllSpecialties() ([]entity.Specialties, error) {

	sqlStmt := `SELECT * FROM specialties`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, sqlStmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specialties []entity.Specialties
	for rows.Next() {
		var s entity.Specialties
		err = rows.Scan(
			&s.ID,
			&s.Name,
			&s.Popular,
			&s.CreatedAt,
			&s.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		specialties = append(specialties, s)
	}

	return specialties, nil
}

func (r *specialtiesRepository) Save(s entity.Specialties) (entity.Specialties, error) {
	sqlStmt := `INSERT INTO specialties (name, popular) VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	_, err := r.db.ExecContext(ctx, sqlStmt, s.Name, s.Popular)

	if err != nil {
		return s, err
	}

	sqlStmt = `SELECT id, name, popular FROM specialties ORDER BY id DESC LIMIT 1`

	row := r.db.QueryRowContext(ctx, sqlStmt)
	err = row.Scan(
		&s.ID,
		&s.Name,
		&s.Popular,
	)

	if err != nil {
		return s, err
	}

	return s, nil
}
