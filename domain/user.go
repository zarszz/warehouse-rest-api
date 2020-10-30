package domain

import (
	"context"
	"time"
)

type User struct {
	ID          string      `json:"id"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Gender      string      `json:"gender"`
	DateOfBirth string      `json:"date_of_birth"`
	Address     UserAddress `json:"user_address"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type UserUseCase interface {
	Fetch(ctx context.Context, num int64) ([]User, string, error)
	GetByID(ctx context.Context, userID string) (User, error)
	Update(ctx context.Context, user *User) error
	Store(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID string) error
}

type UserRepository interface {
	Fetch(ctx context.Context, num int64) (res []User, nextCursor string, err error)
	GetByID(ctx context.Context, userID string) (User, error)
	Update(ctx context.Context, user *User) error
	Store(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID string) error
}
