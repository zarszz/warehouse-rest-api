package models

import (
	"time"

	"gorm.io/gorm"
)

// User - user model struct
type User struct {
	gorm.Model
	IsAdmin     bool
	Email       string
	Password    string
	FirstName   string
	LastName    string
	Gender      string
	DateOfBirth time.Time
	UserAddress UserAddress `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Items       []Item      `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`
}
