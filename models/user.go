package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"column:id" json:"id" gorm:"primary_key;auto_increment;not_null"`
	Username  string    `gorm:"column:username" json:"username"`
	Fullname  string    `gorm:"column:fullname" json:"fullname"`
	Gender    string    `gorm:"column:gender" json:"gender"`
	Email     string    `gorm:"column:email" json:"email"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	CreatedAt time.Time `gorm:"column:created_at" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
