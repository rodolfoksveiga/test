package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string `gorm:"size:255;not null;unique" json:"email" binding:"required"`
	Password string `gorm:"size:255;not null;unique" json:"password" binding:"required"`
}
