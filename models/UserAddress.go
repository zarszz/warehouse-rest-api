package models

import (
	"gorm.io/gorm"
)

// UserAddress - User Address Table
type UserAddress struct {
	gorm.Model
	UserID     uint `gorm:"index"`
	Address    string
	City       string
	State      string
	Country    string
	PostalCode string
}
