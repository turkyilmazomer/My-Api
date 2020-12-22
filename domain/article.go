package domain

import (
	"context"
)

// Article ...
type Article struct {
	ID        int64  `json:"id"`
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Author    Author `json:"author"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	//##BAK
}

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Article, string, error)
	GetByID(ctx context.Context, id int64) (Article, error)
	Update(ctx context.Context, ar *Article) error
	GetByTitle(ctx context.Context, title string) (Article, error)
	Store(context.Context, *Article) error
	Delete(ctx context.Context, id int64) error
}

// ArticleRepository represent the article's repository contract
type ArticleRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Article, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Article, error)
	GetByTitle(ctx context.Context, title string) (Article, error)
	Update(ctx context.Context, ar *Article) error
	Store(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id int64) error
}
