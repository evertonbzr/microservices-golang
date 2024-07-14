package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `gorm:"column:title;type:varchar(255);index"`
	Description string `gorm:"column:description;type:text"`
	Done        bool   `gorm:"column:done;type:boolean;index"`
	UserID      uint   `gorm:"column:user_id;type:integer;index"`
}
