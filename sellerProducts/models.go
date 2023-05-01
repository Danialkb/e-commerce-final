package sellerProducts

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title     string         `gorm:"size: 150;unique"`
	Body      string         `json:"body"`
	MainImage string         `json:"main_image"`
	IsActive  bool           `json:"is_active"`
	IsTop     bool           `json:"is_top"`
	Data      map[string]int `json:"data" gorm:"type:jsonb"`
}
