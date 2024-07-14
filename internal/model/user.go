package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string    `gorm:"column:username;type:varchar(255);index"`
	Email           string    `gorm:"column:email;type:varchar(255);index"`
	Password        string    `gorm:"column:password;type:varchar(255)"`
	IsLogInEnabled  bool      `gorm:"column:is_login_enabled;type:boolean;default:true"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at;type:timestamp"`
}
