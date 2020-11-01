package domain

import (
	"context"
	"time"
)

type Rack struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RackDetail struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	Name      string    `json:"name"`
	Items     []Item    `json:"items"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RackUseCase interface {
	Fetch(ctx context.Context, num int64) (res []Rack, nextCursor string, err error)
	GetByID(ctx context.Context, rackID string) (Rack, error)
	Update(ctx context.Context, rack *Rack) error
	Store(ctx context.Context, rack *Rack) error
	Delete(ctx context.Context, rackID string) error
}

type RackRepository interface {
	Fetch(ctx context.Context, num int64) ([]Rack, string, error)
	GetByID(ctx context.Context, rackID string) (Rack, error)
	Update(ctx context.Context, rack *Rack) error
	Store(ctx context.Context, rack *Rack) error
	Delete(ctx context.Context, rackID string) error
}
