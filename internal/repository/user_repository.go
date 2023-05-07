package repository

import (
	"context"

	"github.com/erikrios/blog/internal/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user entity.User) (id int64, err error)
	FindAll(ctx context.Context) (users []entity.User, err error)
	FindByID(ctx context.Context, id int64) (user entity.User, err error)
	Update(ctx context.Context, id int64, user entity.User) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
