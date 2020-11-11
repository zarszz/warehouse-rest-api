package usacase

import (
	"context"
	"time"

	"github.com/zarszz/warehouse-rest-api/domain"
)

type warehouseUsecase struct {
	warehouseRepo  domain.WarehouseRepository
	contextTimeout time.Duration
}

func NewWarehouseUsecase(warehouseRepository domain.WarehouseRepository, timeout time.Duration) domain.WarehouseUseCase {
	return &warehouseUsecase{
		warehouseRepo:  warehouseRepository,
		contextTimeout: timeout,
	}
}

func (w *warehouseUsecase) Fetch(ctx context.Context, num int64) (warehouses []domain.Warehouse, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()
	return w.warehouseRepo.Fetch(ctx, num)
}

func (w *warehouseUsecase) FetchRoom(ctx context.Context, warehouseID string) (warehouseDetail domain.WarehouseRoom, err error) {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()
	return w.warehouseRepo.FetchRoom(ctx, warehouseID)
}

func (w *warehouseUsecase) GetByID(ctx context.Context, warehouseID string) (warehouse domain.Warehouse, err error) {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()
	return w.warehouseRepo.GetByID(ctx, warehouseID)
}

func (w *warehouseUsecase) Update(ctx context.Context, warehouse *domain.Warehouse) error {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()
	return w.warehouseRepo.Update(ctx, warehouse)
}

func (w *warehouseUsecase) Store(ctx context.Context, warehouse *domain.Warehouse) (id string, err error) {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()
	return w.warehouseRepo.Store(ctx, warehouse)
}

func (w *warehouseUsecase) Delete(ctx context.Context, warehouseID string) error {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()
	return w.warehouseRepo.Delete(ctx, warehouseID)
}
