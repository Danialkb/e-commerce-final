package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null;size:30"`
	LastName  string `gorm:"not null;size:30"`
	Password  string `gorm:"not null;size:255"`
	Email     string `gorm:"not null;size:255;unique"`
	Phone     string `gorm:"not null;size:255;unique"`
}
