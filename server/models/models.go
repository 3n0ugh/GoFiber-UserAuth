package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `json:"username" gorm:"uniqueIndex" validate:"required,min=5,max=12,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
