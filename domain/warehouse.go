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

type WarehouseIndependence struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WarehouseRoom struct {
	ID      string                   `json:"id"`
	Name    string                   `json:"name"`
	Address WarehouseAddress         `json:"address"`
	Rooms   []RoomBelongsToWarehouse `json:"rooms"`
}

type WarehouseDetail struct {
	ID      string           `json:"id"`
	Name    string           `json:"name"`
	Address WarehouseAddress `json:"address"`
	Rooms   []RoomDetail     `json:"rooms"`
}

type WarehouseUseCase interface {
	Fetch(ctx context.Context, num int64) (res []Warehouse, nextCursor string, err error)
	FetchRoom(ctx context.Context, warehouseID string) (res WarehouseRoom, err error)
	GetByID(ctx context.Context, warehouseID string) (Warehouse, error)
	Update(ctx context.Context, warehouse *Warehouse) error
	Store(ctx context.Context, warehouse *Warehouse) (id string, err error)
	Delete(ctx context.Context, warehouseID string) error
	IsWarehouseExist(ctx context.Context, warehouseID string) (bool, error)
}

type WarehouseRepository interface {
	Fetch(ctx context.Context, num int64) ([]Warehouse, string, error)
	FetchRoom(ctx context.Context, warehouseID string) (WarehouseRoom, error)
	GetByID(ctx context.Context, warehouseID string) (Warehouse, error)
	Update(ctx context.Context, warehouse *Warehouse) error
	Store(ctx context.Context, warehouse *Warehouse) (id string, err error)
	Delete(ctx context.Context, warehouseID string) error
	IsWarehouseExist(ctx context.Context, warehouseID string) (bool, error)
}
