package models

import (
	"gorm.io/gorm"
	"time"
)

type Movement struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Value     float64        `json:"value"`
	PocketID  uint           `json:"pocket_id"`
	Pocket    Pocket         `json:"pocket" gorm:"references:PocketID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
