package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService UserServiceInterface
}

func NewUserController() UserController {
	return UserController{userService: NewUserService()}
}

func (u UserController) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if err := u.userService.CreateUser(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"user": user})
		}
	}
}
