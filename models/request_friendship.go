package models

import (
	"time"
)

type Request_friendship struct {
	ID               int       `gorm:"column:id" json:"id" gorm:"primary_key;auto_increment;not_null"`
	Id_user_req_from int       `gorm:"column:id_user_req_from" json:"id_user_req_from"`
	Id_user_req_to   int       `gorm:"column:id_user_req_to" json:"id_user_req_to"`
	Status           int       `gorm:"column:status" json:"status"` //1=request, 2=cancel_request
	CreatedAt        time.Time `gorm:"column:created_at" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
