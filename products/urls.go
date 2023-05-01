package products

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(pc *ProductController, cc *CategoryController) *gin.Engine {
	r := gin.Default()

	// Products endpoints
	r.POST("/products", pc.CreateProduct)
	r.GET("/products/:id", pc.GetProductByID)
	r.PUT("/products", pc.UpdateProduct)
	r.DELETE("/products/:id", pc.DeleteProduct)

	// Categories endpoints
	r.POST("/categories", cc.CreateCategory)
	r.GET("/categories/:id", cc.GetCategoryByID)
	r.PUT("/categories", cc.UpdateCategory)
	r.DELETE("/categories/:id", cc.DeleteCategory)

	return r
}

// How to use? -> for Danial
//func main() {
//	r := gin.Default()
//
//	productService := products.NewProductService()
//	productController := products.NewProductController(productService)
//
//	categoryService := products.NewCategoryService()
//	categoryController := products.NewCategoryController(categoryService)
//
//	products.RegisterRoutes(r, productController, categoryController)
//
//	err := r.Run(":8080")
//	if err != nil {
//		panic(err)
//	}
//}
