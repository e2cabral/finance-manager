package models

type Pocket struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
