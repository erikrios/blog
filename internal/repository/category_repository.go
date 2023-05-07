package repository

import (
	"context"

	"github.com/erikrios/blog/internal/entity"
)

type CategoryRepository interface {
	Insert(ctx context.Context, category entity.Category) (id int64, err error)
	FindAll(ctx context.Context) (categories []entity.Category, err error)
	FindByID(ctx context.Context, id int64) (category entity.Category, err error)
	Update(ctx context.Context, id int64, category entity.Category) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
