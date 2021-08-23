package models

import (
	"gorm.io/gorm"
	"time"
)

type Pocket struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id"`
	User        User           `json:"user" gorm:"references:ID;foreignKey:UserID"`
	Movements   []Movement     `json:"movements" gorm:"references:ID"`
	Type        string         `json:"type"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
