package ratings

import (
	"e-commerce-final/settings"
	"gorm.io/gorm"
	"log"
)

type RatingRepositoryInterface interface {
	CreateRating(*Rating) error
}

type RatingRepositoryV1 struct {
	DB *gorm.DB
}

func NewRatingRepository() *RatingRepositoryV1 {
	db, err := settings.DbSetup()
	if err != nil {
		log.Fatal(err)
		return &RatingRepositoryV1{}
	}
	return &RatingRepositoryV1{DB: db}
}

func (r *RatingRepositoryV1) CreateRating(rating *Rating) error {
	return r.DB.Create(rating).Error
}
