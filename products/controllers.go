package products

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService ProductServiceInterface
	//userService    users.UserServiceInterface
}

type CategoryController struct {
	categoryService CategoryServiceInterface
}

func NewProductController() *ProductController {
	return &ProductController{productService: NewProductService()}
}

func NewCategoryController() *CategoryController {
	return &CategoryController{categoryService: NewCategoryService()}
}

func (pc *ProductController) GetProductAverageRating(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
	}

	avgValue, err := pc.productService.GetProductAverageRating(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	}

	c.JSON(http.StatusOK, gin.H{"data": avgValue})
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	//productList := make([]Product, 0)

	products, err := pc.productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Products not found"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) SearchByName(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name parameter is missing"})
		return
	}
	products, err := pc.productService.SearchByName(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search for products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc ProductController) SearchByPriceRange(c *gin.Context) {
	minPrice, err := strconv.ParseFloat(c.Query("min_price"), 64)
	if err != nil {
		minPrice = 0.0
	}
	maxPrice, err := strconv.ParseFloat(c.Query("max_price"), 64)
	if err != nil {
		maxPrice = math.MaxFloat64
	}
	products, err := pc.productService.SearchByPriceRange(minPrice, maxPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search for products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc ProductController) GetCommentsByProductId(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	commentList, _ := pc.productService.GetCommentsByProductId(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": commentList})
}

func (pc ProductController) CreateProduct(c *gin.Context) {
	//authHeader := c.GetHeader("Authorization")
	//if authHeader == "" {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
	//	return
	//}
	//
	//tokenString := authHeader[7:]
	//
	//token, err := jwt.ParseWithClaims(tokenString, &users.SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(os.Getenv("SECRET_KEY")), nil
	//})
	//fmt.Println(err.Error())
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
	//	return
	//}
	//
	//claims, ok := token.Claims.(*users.SignedDetails)
	//if !ok || !token.Valid {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
	//	return
	//}
	//userType := claims.UserType
	//userID := claims.Id
	//
	//fmt.Println(userType)
	//
	//if userType != "Seller" {
	//	c.JSON(http.StatusForbidden, gin.H{"error": "user is not a seller"})
	//	return
	//}

	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//product.UserID = userID

	if err := pc.productService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": product})
}

func (pc ProductController) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := pc.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (pc ProductController) UpdateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.productService.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (pc ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := pc.productService.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (cc CategoryController) CreateCategory(c *gin.Context) {
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.categoryService.CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": category})
}

func (cc CategoryController) GetCategoryByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := cc.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func (cc CategoryController) UpdateCategory(c *gin.Context) {
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.categoryService.UpdateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func (cc CategoryController) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := cc.categoryService.DeleteCategory(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (cc CategoryController) GetCategories(c *gin.Context) {
	category_list, _ := cc.categoryService.GetCategories()

	c.JSON(http.StatusOK, category_list)
}
