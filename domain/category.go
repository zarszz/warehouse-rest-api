package domain

import (
	"context"
	"time"
)

type Category struct {
	ID        int64     `json:"id"`
	Category  string    `json:"category" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoryUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Category, string, error)
	GetByID(ctx context.Context, id int64) (Category, error)
	Update(ctx context.Context, category *Category) error
	Store(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id int64) error
}

type CategoryRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Category, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Category, error)
	Update(ctx context.Context, category *Category) error
	Store(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id int64) error
}
