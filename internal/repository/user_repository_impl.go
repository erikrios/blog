package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erikrios/blog/internal/entity"
	"github.com/lib/pq"
)

const (
	insertUserQuery = `
INSERT INTO users(username, name, password)
VALUES ($1, $2, $3) RETURNING id;
	`
	findAllUsersQuery = `
SELECT 
id,
username,
name,
password,
created_at,
updated_at,
deleted_at
FROM users
WHERE deleted_at IS NULL
ORDER BY id;
`
	findUserByIDQuery = `
SELECT 
id,
username,
name,
password,
created_at,
updated_at,
deleted_at
FROM users
WHERE id = $1
AND deleted_at IS NULL;
`
	updateUserQuery = `
UPDATE users 
SET username = $2,
name = $3, 
password = $4
WHERE id = $1
AND deleted_at IS NULL;
`
	deleteUserQuery = `
UPDATE users 
SET deleted_at = current_timestamp
WHERE id = $1
AND deleted_at IS NULL;
`
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

func (u *userRepositoryImpl) Insert(ctx context.Context, user entity.User) (id int64, err error) {
	row := u.db.QueryRowContext(ctx, insertUserQuery, user.Username, user.Name, user.Password)

	switch dbErr := row.Scan(&id); dbErr {
	case sql.ErrNoRows:
		err = fmt.Errorf("%w: %w", ErrAlreadyExists, dbErr)
		return
	case nil:
		return
	default:
		if e, ok := dbErr.(*pq.Error); ok {
			if e.Code == "23505" {
				err = fmt.Errorf("%w: %w", ErrAlreadyExists, dbErr)
				return
			}
		}
		err = fmt.Errorf("%w: %w", ErrUnknown, dbErr)
		return
	}
}

func (u *userRepositoryImpl) FindAll(ctx context.Context) (users []entity.User, err error) {
	rows, dbErr := u.db.QueryContext(ctx, findAllUsersQuery)
	if dbErr != nil {
		err = fmt.Errorf("%w: %w", ErrUnknown, dbErr)
		return
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	users = make([]entity.User, 0)
	for rows.Next() {
		var user entity.User
		if dbErr := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Name,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		); dbErr != nil {
			err = fmt.Errorf("%w: %w", ErrUnknown, dbErr)
			return
		}

		users = append(users, user)
	}

	return
}

func (u *userRepositoryImpl) FindByID(ctx context.Context, id int64) (user entity.User, err error) {
	row := u.db.QueryRowContext(ctx, findUserByIDQuery, id)

	switch dbErr := row.Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	); dbErr {
	case sql.ErrNoRows:
		err = fmt.Errorf("%w: %w", ErrNotFound, dbErr)
		return
	case nil:
		return
	default:
		err = fmt.Errorf("%w: %w", ErrUnknown, dbErr)
		return
	}
}

func (u *userRepositoryImpl) Update(ctx context.Context, id int64, user entity.User) (err error) {
	res, dbErr := u.db.ExecContext(ctx, updateUserQuery, id, user.Username, user.Name, user.Password)
	if dbErr != nil {
		err = fmt.Errorf("%w: %w", ErrUnknown, dbErr)
		return
	}

	n, dbErr := res.RowsAffected()
	if dbErr != nil {
		err = fmt.Errorf("%w: %w", ErrUnknown, dbErr)
		return
	}

	if n < 1 {
		err = fmt.Errorf("%w: %w", ErrUnknown, errors.New("count: rows affected less than 1"))
		return
	}

	return
}

func (u *userRepositoryImpl) Delete(ctx context.Context, id int64) (err error) {
	res, dbErr := u.db.ExecContext(ctx, deleteUserQuery, id)
	if dbErr != nil {
		err = fmt.Errorf("%w: %w", ErrUnknown, dbErr)
		return
	}

	n, dbErr := res.RowsAffected()
	if dbErr != nil {
		err = fmt.Errorf("%w: %w", ErrUnknown, dbErr)
		return
	}

	if n < 1 {
		err = fmt.Errorf("%w: %w", ErrUnknown, errors.New("count: rows affected less than 1"))
		return
	}

	return
}
