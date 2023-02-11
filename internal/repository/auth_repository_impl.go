package repository

import (
	"context"
	"time"

	"github.com/muhangga/internal/entity"
)

var timeOut = 15 * time.Second

func (r *repository) Save(u entity.User) (entity.User, error) {
	var sqlStmt string

	sqlStmt = "INSERT INTO public.users (full_name, email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	_, err := r.db.ExecContext(ctx, sqlStmt, u.FullName, u.Email, u.Password, u.Role, u.CreatedAt, u.UpdatedAt)

	if err != nil {
		return u, err
	}

	sqlStmt = "SELECT id, full_name, email, role, created_at, updated_at FROM users ORDER BY id DESC LIMIT 1"

	row := r.db.QueryRowContext(ctx, sqlStmt)
	err = row.Scan(
		&u.ID,
		&u.FullName,
		&u.Email,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	queryStmt := `SELECT email FROM public.users WHERE email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	err := r.db.QueryRowContext(ctx, queryStmt, email).Scan(
		&user.Email,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(id int64) (entity.User, error) {
	var user entity.User

	queryStmt := `SELECT * FROM public.users WHERE id = $1`

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
	); err != nil {
		return user, err
	}

	return user, nil

}
