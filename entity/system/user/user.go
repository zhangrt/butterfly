package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"size:128"`
	Account string `gorm:"size:20"`
	Password string `gorm:"size:256"`
	Gender string `gorm:"size:1"`
	Email string `gorm:"size:100"`
	Telephone string `gorm:"size:20"`
}
