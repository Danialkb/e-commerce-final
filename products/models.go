package products

import (
	"e-commerce-final/users"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	User     users.User `gorm:"foreignKey:user"`
	Category Category   `gorm:"foreignKey:category"`
	//CategoryID uint
	Title string `gorm:"size: 150;unique"`

	Body       string `json:"body"`
	MainImage  string `json:"main_image"`
	IsActive   bool   `json:"is_active"`
	IsTop      bool   `json:"is_top"`
	UserID     uint   `gorm:"column:user_id"`
	CategoryID uint   `gorm:"column:category_id"`
}

type Category struct {
	gorm.Model
	Title string `gorm:"size: 150; unique"`
}
