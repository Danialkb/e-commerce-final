package products

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title      string `gorm:"size: 150;unique"`
	Body       string `json:"body"`
	MainImage  string `json:"main_image"`
	Price      float64
	IsActive   bool `json:"is_active"`
	IsTop      bool `json:"is_top"`
	CategoryID uint `json:"category_id" binding:"required" gorm:"ForeignKey:Category.ID"`
	UserID     uint `json:"user_id" binding:"required" gorm:"ForeignKey:User.ID"`
}

type Category struct {
	gorm.Model
	Title string `gorm:"size: 150; unique"`
}
