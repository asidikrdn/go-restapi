package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MstUser struct {
	ID              uuid.UUID `gorm:"type:varchar(255);primaryKey"`
	FullName        string    `gorm:"type:varchar(255)"`
	Email           string    `gorm:"type:varchar(255);unique"`
	IsEmailVerified bool
	Phone           string `gorm:"type:varchar(255);unique"`
	IsPhoneVerified bool
	Address         string `gorm:"type:text"`
	Password        string `gorm:"type:varchar(255)"`
	RoleID          uint
	Image           string `gorm:"type:varchar(255)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	gorm.DeletedAt
	// parent
	Role MstRole
}
