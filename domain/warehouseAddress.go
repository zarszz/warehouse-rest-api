package domain

import (
	"context"
	"time"
)

type WarehouseAddress struct {
	AddressID   string    `json:"address_id"`
	WarehouseID string    `json:"warehouse_id"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	PostalCode  string    `json:"postal_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WarehouseAddressUsecase interface {
	FetchWarehouseAddress(ctx context.Context, warehouseID string) (*WarehouseAddress, error)
	Update(ctx context.Context, warehouseAddress WarehouseAddress) error
	Store(ctx context.Context, warehouseAddress WarehouseAddress) error
}

type WarehouseAddressRepository interface {
	FetchWarehouseAddress(ctx context.Context, warehouseID string) (res *WarehouseAddress, err error)
	Update(ctx context.Context, warehouseAddress WarehouseAddress) (err error)
	Store(ctx context.Context, warehouseAddress WarehouseAddress) (err error)
}
