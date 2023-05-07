package repository

import (
	"context"

	"github.com/erikrios/blog/internal/entity"
)

type ArticleRepository interface {
	Insert(ctx context.Context, article entity.Article) (id int64, err error)
	FindAll(ctx context.Context) (articles []entity.Article, err error)
	FindByID(ctx context.Context, id int64) (article entity.Article, err error)
	Update(ctx context.Context, id int64, article entity.Article) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
