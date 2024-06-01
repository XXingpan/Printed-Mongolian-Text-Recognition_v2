package login

import (
	"backend/config"
	"backend/dao/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	db := config.DB
	userService := service.UserService{DB: db}

	var registerData struct {
		Username string `form:"username" binding:"required,min=3,max=20"`
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required,min=6"`
	}

	if err := c.Bind(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	if exists, err := userService.UserExists(registerData.Username, registerData.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while checking for existing user"})
		return
	} else if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
		return
	}

	if err := userService.CreateUser(registerData.Username, registerData.Email, registerData.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	userID, err := userService.GetUserByUsernameAndPassword(registerData.Username, registerData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while getting user's id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful", "userID": userID})
}
