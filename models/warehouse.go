package models

import (
	"time"

	"gorm.io/gorm"
)

// Warehouse - warehouse model struct
type Warehouse struct {
	ID            string `gorm:"primaryKey"`
	WarehouseName string
	Items         []Item           `gorm:"foreignKey:ItemCode;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE"`
	Address       WarehouseAddress `gorm:"foreignKey:WarehouseID;constraint:OnUpdate:NO ACTION,OnDelete:CASCADE"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
