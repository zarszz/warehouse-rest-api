package domain

import (
	"context"
	"time"
)

type Room struct {
	ID          string    `json:"id"`
	WarehouseID string    `json:"warehouse_id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RoomDetail struct {
	ID          string       `json:"id"`
	WarehouseID string       `json:"warehouse_id"`
	Racks       []RackDetail `json:"racks"`
	Name        string       `json:"name"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type RoomBelongsToWarehouse struct {
	ID            string    `json:"id"`
	WarehouseID   string    `json:"warehouse_id"`
	Name          string    `json:"name"`
	WarehouseName string    `json:"warehouse_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RoomUseCase interface {
	Fetch(ctx context.Context, num int64) (res []Room, nextCursor string, err error)
	GetByID(ctx context.Context, roomID string) (Room, error)
	Update(ctx context.Context, room *Room) error
	Store(ctx context.Context, room *Room) error
	Delete(ctx context.Context, roomID string) error
}

type RoomRepository interface {
	Fetch(ctx context.Context, num int64) ([]Room, string, error)
	GetByID(ctx context.Context, roomID string) (Room, error)
	Update(ctx context.Context, room *Room) error
	Store(ctx context.Context, room *Room) error
	Delete(ctx context.Context, roomID string) error
}
