package domain

import (
	"context"
	"time"
)

type UserAddress struct {
	AddressID  string    `json:"address_id"`
	UserID     string    `json:"user_id"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	Country    string    `json:"country"`
	PostalCode string    `json:"postal_code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserAddressUsecase interface {
	FetchUserAddress(ctx context.Context, userID string) (*UserAddress, error)
	Update(ctx context.Context, userAddress UserAddress) error
	Store(ctx context.Context, userAddress UserAddress) error
}

type UserAddressRepository interface {
	FetchUserAddress(ctx context.Context, userID string) (res *UserAddress, err error)
	Update(ctx context.Context, userAddress UserAddress) (err error)
	Store(ctx context.Context, userAddress UserAddress) (err error)
}
