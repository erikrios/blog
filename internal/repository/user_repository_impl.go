package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/erikrios/blog/internal/entity"
	"github.com/lib/pq"
)

const (
	insertUserQuery = `
INSERT INTO users(username, name, password)
VALUES ($1, $2, $3) RETURNING id;
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

func (u *userRepositoryImpl) FindAll(ctx context.Context) (users []entity.User, err error) { return }

func (u *userRepositoryImpl) FindByID(ctx context.Context, id int64) (user entity.User, err error) {
	return
}

func (u *userRepositoryImpl) Update(ctx context.Context, id int64, user entity.User) (err error) {
	return
}

func (u *userRepositoryImpl) Delete(ctx context.Context, id int64) (err error) { return }
