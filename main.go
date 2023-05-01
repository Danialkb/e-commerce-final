package main

import (
	"e-commerce-final/settings"
	"e-commerce-final/users"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, _ := settings.DbSetup()
	err := db.AutoMigrate(&users.User{})
	if err != nil {
		log.Fatalf("Error %s", err)
	}
}
