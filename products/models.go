package products

import (
	"e-commerce-final/users"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	User      users.User `gorm:"foreignKey:id"`
	Title     string     `gorm:"size: 150;unique"`
	Body      string     `json:"body"`
	MainImage string     `json:"main_image"`
	IsActive  bool       `json:"is_active"`
	IsTop     bool       `json:"is_top"`
}
