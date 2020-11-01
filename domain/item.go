package domain

import (
	"context"
	"time"
)

type Item struct {
	ItemID      string    `json:"item_id"`
	CategoryID  string    `json:"category_id"`
	WarehouseID string    `json:"warehouse_id"`
	OwnerID     string    `json:"owner_id"`
	RackID      string    `json:"rack_id"`
	ItemName    string    `json:"item_name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ItemUseCase interface {
	Fetch(ctx context.Context, num int64) ([]Item, string, error)
	GetByID(ctx context.Context, itemID string) (Item, error)
	Update(ctx context.Context, item *Item) error
	Store(ctx context.Context, item *Item) error
	Delete(ctx context.Context, itemID string) error
}

type ItemRepository interface {
	Fetch(ctx context.Context, num int64) (res []Item, nextCursor string, err error)
	GetByID(ctx context.Context, itemID string) (Item, error)
	Update(ctx context.Context, item *Item) error
	Store(ctx context.Context, item *Item) error
	Delete(ctx context.Context, itemID string) error
}
