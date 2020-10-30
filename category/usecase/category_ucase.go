package usecase

import (
	"context"
	"time"

	"github.com/zarszz/warehouse-rest-api/domain"
)

type categoryUsecase struct {
	categoryRepo   domain.CategoryRepository
	contextTimeout time.Duration
}

// NewCategoryUsecase will create new an categoryUsecase object representation of domain.CategoryUsecase interface
func NewCategoryUsecase(categoryRepository domain.CategoryRepository, timeout time.Duration) domain.CategoryRepository {
	return &categoryUsecase{
		categoryRepo:   categoryRepository,
		contextTimeout: timeout,
	}
}

func (a *categoryUsecase) Fetch(c context.Context, cursor string, num int64) (res []domain.Category, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, nextCursor, err = a.categoryRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}
	return
}

func (a *categoryUsecase) GetByID(c context.Context, id string) (res domain.Category, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	return
}

func (a *categoryUsecase) Update(c context.Context, category *domain.Category) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	category.UpdatedAt = time.Now()
	return a.categoryRepo.Update(ctx, category)
}

func (a *categoryUsecase) Store(c context.Context, category *domain.Category) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	err = a.categoryRepo.Store(ctx, category)
	return
}

func (a *categoryUsecase) Delete(c context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedCategory, err := a.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedCategory == (domain.Category{}) {
		return domain.ErrNotFound
	}
	return a.categoryRepo.Delete(ctx, id)
}
