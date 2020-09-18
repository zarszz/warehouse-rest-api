package models

import (
	"gorm.io/gorm"
)

// Category - category model struct
type Category struct {
	gorm.Model
	Category string
}
