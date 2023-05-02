package main

import (
	"e-commerce-final/products"
	"e-commerce-final/settings"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	pc := products.NewProductController()
	cc := products.NewCategoryController()

	db, err := settings.DbSetup()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&products.Product{}, &products.Category{})

	r.POST("/products", pc.CreateProduct)
	r.GET("/products/:id", pc.GetProductByID)
	r.PUT("/products", pc.UpdateProduct)
	r.DELETE("/products/:id", pc.DeleteProduct)

	// Categories endpoints
	r.POST("/categories", cc.CreateCategory)
	r.GET("/categories/:id", cc.GetCategoryByID)
	r.PUT("/categories", cc.UpdateCategory)
	r.DELETE("/categories/:id", cc.DeleteCategory)

	log.Fatal(http.ListenAndServe(":8080", r))
}
