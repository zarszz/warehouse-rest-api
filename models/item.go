package models

import (
	"time"

	"gorm.io/gorm"
)

// Item - item model struct
type Item struct {
	ItemCode    string `gorm:"primaryKey"`
	ItemName    string
	CategoryID  uint
	WarehouseID string
	OwnerID     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
