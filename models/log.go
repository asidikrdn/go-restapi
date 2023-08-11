package models

import (
	"time"
)

type Log struct {
	ID          uint
	Date        time.Time
	IPAddress   string `gorm:"type:varchar(255)"`
	Host        string `gorm:"type:varchar(255)"`
	Path        string `gorm:"type:varchar(255)"`
	Method      string `gorm:"type:varchar(255)"`
	Body        string
	File        string
	ResposeTime float64
	StatusCode  int
	// Message     string `gorm:"type:varchar(255)"`
}
