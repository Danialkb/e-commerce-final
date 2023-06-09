package main

import (
	_ "github.com/lib/pq"
)

<<<<<<< HEAD
//
//func main() {
//	router := gin.Default()
//	userController := users.NewUserController()
//	pc := products.NewProductController()
//	cc := products.NewCategoryController()
//	oc := orders.NewOrderController()
//
//	db, _ := settings.DbSetup()
//	db.AutoMigrate(&users.User{}, &products.Product{}, &products.Category{}, &orders.Order{})
//
//	router.POST("/users", userController.CreateUser)
//	router.GET("/users/:user_id", userController.GetUserById)
//	router.POST("/users/token", userController.LogIn)
//
//	router.POST("/orders", oc.CreateOrder)
//	router.GET("/orders/:id", oc.GetOrderById)
//	router.PUT("/orders/:id", oc.UpdateOrder)
//	router.GET("/users/:user_id/orders", oc.GetOrders)
//	router.DELETE("/orders/:id", oc.DeleteOrder)
//
//	router.POST("/products", pc.CreateProduct)
//	router.GET("/products/:id", pc.GetProductByID)
//	router.PUT("/products", pc.UpdateProduct)
//	router.DELETE("/products/:id", pc.DeleteProduct)
//
//	// Categories endpoints
//	router.POST("/categories", cc.CreateCategory)
//	router.GET("/categories/:id", cc.GetCategoryByID)
//	router.PUT("/categories", cc.UpdateCategory)
//	router.DELETE("/categories/:id", cc.DeleteCategory)
//
//	log.Fatal(http.ListenAndServe(":8080", router))
//	//t, b, _ := users.TokenGenerator("bidaybek044@gmail.com", "Danial", "Bidaibek", "1234")
//}
=======
func main() {
	router := gin.Default()
	userController := users.NewUserController()
	pc := products.NewProductController()
	cc := products.NewCategoryController()
	oc := orders.NewOrderController()

	db, _ := settings.DbSetup()
	db.AutoMigrate(&users.User{}, &products.Product{}, &products.Category{}, &orders.Order{})

	router.POST("/users", userController.CreateUser)
	router.GET("/users/:user_id", userController.GetUserById)
	router.POST("/users/token", userController.LogIn)

	router.POST("/orders", oc.CreateOrder)
	router.GET("/orders/:id", oc.GetOrderById)
	router.PUT("/orders/:id", oc.UpdateOrder)
	router.GET("/users/:user_id/orders", oc.GetOrders)
	router.DELETE("/orders/:id", oc.DeleteOrder)

	router.POST("/products", pc.CreateProduct)
	router.GET("/products/", pc.SearchByName)
	router.GET("/products/filter", pc.SearchByPriceRange)
	router.GET("/products/:id", pc.GetProductByID)
	router.PUT("/products", pc.UpdateProduct)
	router.DELETE("/products/:id", pc.DeleteProduct)

	// Categories endpoints
	router.POST("/categories", cc.CreateCategory)
	router.GET("/categories/:id", cc.GetCategoryByID)
	router.PUT("/categories", cc.UpdateCategory)
	router.DELETE("/categories/:id", cc.DeleteCategory)

	log.Fatal(http.ListenAndServe(":8080", router))
	//t, b, _ := users.TokenGenerator("bidaybek044@gmail.com", "Danial", "Bidaibek", "1234")
}
>>>>>>> bbc675b4d7b13b506b02105be17923e4617ff854
