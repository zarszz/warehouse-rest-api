package usecase

import (
	"context"
	"time"

	"github.com/zarszz/warehouse-rest-api/domain"
)

type rackUsecase struct {
	rackRepo       domain.RackRepository
	contextTimeout time.Duration
}

// NewRackUsecase will create new an rackUsecase object representation of domain.RackUsecase interface
func NewRackUsecase(rackRepository domain.RackRepository, timeout time.Duration) domain.RackRepository {
	return &rackUsecase{
		rackRepo:       rackRepository,
		contextTimeout: timeout,
	}
}

func (a *rackUsecase) Fetch(c context.Context, cursor string, num int64) (res []domain.Rack, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, nextCursor, err = a.rackRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}
	return
}

func (a *rackUsecase) GetByID(c context.Context, id string) (res domain.Rack, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.rackRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	return
}

func (a *rackUsecase) Update(c context.Context, category *domain.Rack) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	category.UpdatedAt = time.Now()
	return a.rackRepo.Update(ctx, category)
}

func (a *rackUsecase) Store(c context.Context, rack *domain.Rack) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	err = a.rackRepo.Store(ctx, rack)
	return
}

func (a *rackUsecase) Delete(c context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedRack, err := a.rackRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedRack == (domain.Rack{}) {
		return domain.ErrNotFound
	}
	return a.rackRepo.Delete(ctx, id)
}
