package models

import (
	"time"
)

type Status_friendship struct {
	ID                    int       `gorm:"column:id" json:"id" gorm:"primary_key;auto_increment;not_null"`
	Id_request_friendship int       `gorm:"column:id_request_friendship" json:"id_request_friendship"`
	Status                int       `gorm:"column:status" json:"status"` //1=accept, 2=reject
	CreatedAt             time.Time `gorm:"column:created_at" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt             time.Time `gorm:"column:updated_at" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
