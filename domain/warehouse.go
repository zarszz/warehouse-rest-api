package domain

import (
	"context"
	"time"
)

type Warehouse struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Address   WarehouseAddress `json:"address"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type WarehouseDetail struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Address   WarehouseAddress `json:"address"`
	Rooms     []RoomDetail     `json:"items"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type WarehouseUseCase interface {
	Fetch(ctx context.Context, num int64) (res []Warehouse, nextCursor string, err error)
	GetByID(ctx context.Context, warehouseID string) (Warehouse, error)
	Update(ctx context.Context, warehouse *Warehouse) error
	Store(ctx context.Context, warehouse *Warehouse) error
	Delete(ctx context.Context, warehouseID string) error
}

type WarehouseRepository interface {
	Fetch(ctx context.Context, num int64) ([]Warehouse, string, error)
	GetByID(ctx context.Context, warehouseID string) (Warehouse, error)
	Update(ctx context.Context, warehouse *Warehouse) error
	Store(ctx context.Context, warehouse *Warehouse) error
	Delete(ctx context.Context, warehouseID string) error
}
