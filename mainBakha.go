package main

import (
	"e-commerce-final/comments"
	"e-commerce-final/products"
	"e-commerce-final/ratings"
	"e-commerce-final/settings"
	"e-commerce-final/users"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	pc := products.NewProductController()
	cc := products.NewCategoryController()
	com_c := comments.NewCommentController()
	rc := ratings.NewRatingController()
	db, err := settings.DbSetup()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&products.Product{}, &products.Category{}, &users.User{}, &comments.Comment{}, &ratings.Rating{})

	userController := users.NewUserController()

	//Users endpoints
	r.POST("/users", userController.CreateUser)
	r.POST("/users/token", userController.LogIn)

	//Comments endpoints
	r.POST("/comments", com_c.CreateComment)
	r.GET("/comments/:id", com_c.GetCommentByID)
	r.DELETE("/comments/:id", com_c.DeleteComments)

	//Ratings endpoints
	r.POST("/ratings", rc.CreateRating)

	// Products endpoints
	r.GET("/products", pc.GetProducts)
	r.GET("/products/:id/comments", pc.GetCommentsByProductId)
	r.POST("/products", pc.CreateProduct)
	r.GET("/products/:id", pc.GetProductByID)
	r.PUT("/products", pc.UpdateProduct)
	r.DELETE("/products/:id", pc.DeleteProduct)
	r.GET("/products/:id/rating", pc.GetProductAverageRating)

	// Categories endpoints
	r.GET("/categories", cc.GetCategories)
	r.POST("/categories", cc.CreateCategory)
	r.GET("/categories/:id", cc.GetCategoryByID)
	r.PUT("/categories/:id", cc.UpdateCategory)
	r.DELETE("/categories/:id", cc.DeleteCategory)

	log.Fatal(http.ListenAndServe(":8080", r))
}

//package main
//
//import (
//	"e-commerce-final/products"
//	"e-commerce-final/settings"
//	"e-commerce-final/users"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//)
//
//func main() {
//	r := gin.Default()
//	pc := products.NewProductController()
//	cc := products.NewCategoryController()
//
//	db, err := settings.DbSetup()
//	if err != nil {
//		log.Fatal(err)
//	}
//	db.AutoMigrate(&products.Product{}, &products.Category{}, &users.User{})
//
//	userController := users.NewUserController()
//
//	r.POST("/users", userController.CreateUser)
//	r.POST("/users/token", userController.LogIn)
//
//	r.POST("/products", pc.CreateProduct)
//	r.GET("/products/:id", pc.GetProductByID)
//	r.PUT("/products", pc.UpdateProduct)
//	r.DELETE("/products/:id", pc.DeleteProduct)
//
//	// Categories endpoints
//	r.POST("/categories", cc.CreateCategory)
//	r.GET("/categories/:id", cc.GetCategoryByID)
//	r.PUT("/categories", cc.UpdateCategory)
//	r.DELETE("/categories/:id", cc.DeleteCategory)
//
//	log.Fatal(http.ListenAndServe(":8080", r))
//}
