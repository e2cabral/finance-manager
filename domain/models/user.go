package models

type User struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Pockets  []Pocket `json:"pockets" gorm:"foreignKey:UserID"`
}
