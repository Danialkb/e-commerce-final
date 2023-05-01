package users

import (
	"e-commerce-final/settings"
	"gorm.io/gorm"
	"log"
)

type UserReposInterface interface {
	CreateUser(u *User) error
}

type UserReposV1 struct {
	DB *gorm.DB
}

func NewUserRepos() *UserReposV1 {
	db, err := settings.DbSetup()
	if err != nil {
		log.Fatalf("Error %s", err)
		return &UserReposV1{}
	}
	return &UserReposV1{DB: db}
}

func (u *UserReposV1) CreateUser(user *User) error {
	return u.DB.Create(&user).Error
}
