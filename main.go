package main

import (
	"e-commerce-final/settings"
	"e-commerce-final/users"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, _ := settings.DbSetup()
	err := db.AutoMigrate(&users.User{})
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	t, b, _ := users.TokenGenerator("bidaybek044@gmail.com", "Danial", "Bidaibek", "1234")
	fmt.Println(t)
	fmt.Println(b)
}
