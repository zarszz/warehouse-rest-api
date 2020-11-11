package usacase

import (
	"context"
	"time"

	"github.com/zarszz/warehouse-rest-api/domain"
)

type itemUsecase struct {
	itemRepo       domain.ItemRepository
	contextTimeout time.Duration
}

func NewItemUsecase(itemRepository domain.ItemRepository, timeout time.Duration) domain.ItemUseCase {
	return &itemUsecase{
		itemRepo:       itemRepository,
		contextTimeout: timeout,
	}
}

func (r *itemUsecase) Fetch(ctx context.Context, num int64) (items []domain.Item, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.itemRepo.Fetch(ctx, num)
}
func (r *itemUsecase) GetByID(ctx context.Context, itemID string) (room domain.Item, err error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.itemRepo.GetByID(ctx, itemID)
}

func (r *itemUsecase) GetByRackID(ctx context.Context, roomID string) (rooms []domain.Item, err error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.itemRepo.GetByRackID(ctx, roomID)
}
func (r *itemUsecase) Update(ctx context.Context, item *domain.Item) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.itemRepo.Update(ctx, item)
}
func (r *itemUsecase) Store(ctx context.Context, item *domain.Item) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.itemRepo.Store(ctx, item)
}
func (r *itemUsecase) Delete(ctx context.Context, itemID string) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.itemRepo.Delete(ctx, itemID)
}
