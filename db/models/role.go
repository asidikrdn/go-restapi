package models

import (
	"gorm.io/gorm"
)

// Role
type MstRole struct {
	gorm.Model
	Role string `gorm:"unique"`
}
