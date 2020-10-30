package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/zarszz/warehouse-rest-api/domain"
)

type userAddressUsecase struct {
	userAddressRepository domain.UserAddressRepository
	contextTimeout        time.Duration
}

func NewUserAddressUsecase(userAddressRepository domain.UserAddressRepository, timeout time.Duration) domain.UserAddressUsecase {
	return &userAddressUsecase{
		userAddressRepository: userAddressRepository,
		contextTimeout:        timeout,
	}
}

func (useCase *userAddressUsecase) FetchUserAddress(ctx context.Context, userID string) (*domain.UserAddress, error) {
	if userID == "" {
		return nil, errors.New("User Id Should not empty")
	}
	ctx, cancel := context.WithTimeout(ctx, useCase.contextTimeout)
	defer cancel()

	return useCase.userAddressRepository.FetchUserAddress(ctx, userID)
}

func (useCase *userAddressUsecase) Store(ctx context.Context, userAddress domain.UserAddress) error {
	ctx, cancel := context.WithTimeout(ctx, useCase.contextTimeout)
	defer cancel()

	return useCase.userAddressRepository.Store(ctx, userAddress)
}

func (useCase *userAddressUsecase) Update(ctx context.Context, userAddress domain.UserAddress) error {
	ctx, cancel := context.WithTimeout(ctx, useCase.contextTimeout)
	defer cancel()

	return useCase.userAddressRepository.Update(ctx, userAddress)
}
