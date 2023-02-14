package repository

import (
	"context"

	"github.com/muhangga/internal/entity"
)

func (r *userRepository) GetAllUser() ([]entity.User, error) {
	sqlStmt := `SELECT * FROM users`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var u entity.User
		if err := rows.Scan(
			&u.ID,
			&u.FullName,
			&u.Email,
			&u.Password,
			&u.PhoneNumber,
			&u.Gender,
			&u.Avatar,
			&u.AboutMe,
			&u.Role,
			&u.CreatedAt,
			&u.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *userRepository) UpdateUser(e entity.User) (entity.User, error) {
	sqlStmt := `UPDATE users SET full_name = $1, gender = $2, avatar = $3, email = $4, phone_number = $5, about_me = $6, updated_at = $7 WHERE id = $8`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	_, err := r.db.ExecContext(ctx, sqlStmt, e.FullName, e.Gender, e.Avatar, e.Email, e.PhoneNumber, e.AboutMe, e.UpdatedAt, e.ID)
	if err != nil {
		return e, err
	}

	sqlStmt = `SELECT * FROM users WHERE id = $1`

	row := r.db.QueryRowContext(ctx, sqlStmt, e.ID)
	err = row.Scan(
		&e.ID,
		&e.FullName,
		&e.Email,
		&e.Password,
		&e.PhoneNumber,
		&e.Gender,
		&e.Avatar,
		&e.AboutMe,
		&e.Role,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		return e, err
	}

	return e, nil
}

func (r *userRepository) FindByID(id int64) (entity.User, error) {
	var user entity.User

	queryStmt := `SELECT * FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	if err := r.db.QueryRowContext(ctx, queryStmt, id).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.Gender,
		&user.Avatar,
		&user.AboutMe,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return user, err
	}

	return user, nil

}
