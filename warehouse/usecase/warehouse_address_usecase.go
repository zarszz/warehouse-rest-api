package usacase

import (
	"context"
	"errors"
	"time"

	"github.com/zarszz/warehouse-rest-api/domain"
)

type warehouseAddressUsecase struct {
	warehouseAddressRepo domain.WarehouseAddressRepository
	contextTimeout       time.Duration
}

func NewWarehouseAddressUsecase(warehouseAddressRepository domain.WarehouseAddressRepository, timeout time.Duration) domain.WarehouseAddressUsecase {
	return &warehouseAddressUsecase{
		warehouseAddressRepo: warehouseAddressRepository,
		contextTimeout:       timeout,
	}
}

func (w *warehouseAddressUsecase) FetchWarehouseAddress(ctx context.Context, warehouseID string) (*domain.WarehouseAddress, error) {
	if warehouseID == "" {
		return nil, errors.New("User Id Should not empty")
	}
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

	return w.warehouseAddressRepo.FetchWarehouseAddress(ctx, warehouseID)
}

func (w *warehouseAddressUsecase) Store(ctx context.Context, warehouseAddress domain.WarehouseAddress) error {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

	return w.warehouseAddressRepo.Store(ctx, warehouseAddress)
}

func (w *warehouseAddressUsecase) Update(ctx context.Context, warehouseAddress domain.WarehouseAddress) error {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

	return w.warehouseAddressRepo.Update(ctx, warehouseAddress)
}
