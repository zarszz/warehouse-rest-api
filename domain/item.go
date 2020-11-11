package domain

import (
	"context"
	"time"
)

type Item struct {
	ID          string    `json:"id"`
	ItemName    string    `json:"item_name"`
	Description string    `json:"description"`
	OwnerID     string    `json:"owner_id"`
	CategoryID  string    `json:"category_id"`
	RackID      string    `json:"rack_id"`
	RoomID      string    `json:"room_id"`
	WarehouseID string    `json:"warehouse_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ItemUseCase interface {
	Fetch(ctx context.Context, num int64) ([]Item, string, error)
	GetByID(ctx context.Context, itemID string) (Item, error)
	GetByRackID(ctx context.Context, rackID string) ([]Item, error)
	Update(ctx context.Context, item *Item) error
	Store(ctx context.Context, item *Item) error
	Delete(ctx context.Context, itemID string) error
}

type ItemRepository interface {
	Fetch(ctx context.Context, num int64) (res []Item, nextCursor string, err error)
	GetByID(ctx context.Context, itemID string) (Item, error)
	GetByRackID(ctx context.Context, roomID string) ([]Item, error)
	Update(ctx context.Context, item *Item) error
	Store(ctx context.Context, item *Item) error
	Delete(ctx context.Context, itemID string) error
}
