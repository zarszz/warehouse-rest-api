package usecase

import (
	"context"
	"time"

	"github.com/zarszz/warehouse-rest-api/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &userUsecase{
		userRepo:       userRepository,
		contextTimeout: timeout,
	}
}

func (a *userUsecase) Fetch(ctx context.Context, num int64) (res []domain.User, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	res, nextCursor, err = a.userRepo.Fetch(ctx, num)
	if err != nil {
		return nil, "", err
	}
	return

}
func (a *userUsecase) GetByID(ctx context.Context, userID string) (user domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	user, err = a.userRepo.GetByID(ctx, userID)
	if err != nil {
		return
	}
	return
}

func (a *userUsecase) Update(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	user.UpdatedAt = time.Now()
	return a.userRepo.Update(ctx, user)
}

func (a *userUsecase) Store(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	err = a.userRepo.Store(ctx, user)
	return
}
func (a *userUsecase) Delete(ctx context.Context, userID string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	existedUser, err := a.userRepo.GetByID(ctx, userID)
	if err != nil {
		return
	}
	if existedUser == (domain.User{}) {
		return domain.ErrNotFound
	}
	return a.userRepo.Delete(ctx, userID)
}
