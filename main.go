package main

import (
	"e-commerce-final/users"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	userController := users.NewUserController()

	router.POST("/users", userController.CreateUser)
	router.POST("/users/token", userController.LogIn)

	log.Fatal(http.ListenAndServe(":8080", router))
	//t, b, _ := users.TokenGenerator("bidaybek044@gmail.com", "Danial", "Bidaibek", "1234")
}
