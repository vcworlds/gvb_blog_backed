package models

import "time"

type MODEL struct {
	Id       uint      `json:"id" gorm:"primary_key"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}
