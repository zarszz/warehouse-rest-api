package models

import (
	"gorm.io/gorm"
)

// WarehouseAddress - Warehouse Address Table
type WarehouseAddress struct {
	gorm.Model
	WarehouseID string
	Address     string
	Province    string
	City        string
	Country     string
	PostalCode  string
}
