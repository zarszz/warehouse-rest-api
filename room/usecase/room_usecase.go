package usacase

import (
	"context"
	"time"

	"github.com/zarszz/warehouse-rest-api/domain"
)

type roomUsecase struct {
	roomRepo       domain.RoomRepository
	contextTimeout time.Duration
}

func NewRoomeUsecase(roomRepository domain.RoomRepository, timeout time.Duration) domain.RoomUseCase {
	return &roomUsecase{
		roomRepo:       roomRepository,
		contextTimeout: timeout,
	}
}

func (r *roomUsecase) Fetch(ctx context.Context, num int64) (warehouses []domain.Room, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.roomRepo.Fetch(ctx, num)
}
func (r *roomUsecase) GetByID(ctx context.Context, roomID string) (room domain.Room, err error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.roomRepo.GetByID(ctx, roomID)
}
func (r *roomUsecase) Update(ctx context.Context, room *domain.Room) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.roomRepo.Update(ctx, room)
}
func (r *roomUsecase) Store(ctx context.Context, room *domain.Room) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.roomRepo.Store(ctx, room)
}
func (r *roomUsecase) Delete(ctx context.Context, roomID string) error {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.roomRepo.Delete(ctx, roomID)
}
