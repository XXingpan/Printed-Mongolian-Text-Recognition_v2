package login

import (
	"backend/config"
	"backend/dao/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(c *gin.Context) {
	db := config.DB
	userService := service.UserService{DB: db}
	var loginData struct {
		Username string `form:"username" binding:"required,min=3,max=20"`
		Password string `form:"password" binding:"required,min=6"`
	}

	// 尝试绑定输入的表单数据到loginData结构体
	if err := c.Bind(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request parameters",
		})
		return
	}

	// 使用UserService检查用户名和密码是否正确
	userID, err := userService.GetUserByUsernameAndPassword(loginData.Username, loginData.Password)
	if err != nil {
		// 假设错误意味着用户未找到或密码不匹配
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authentication failed",
		})
		return
	}

	// 如果验证成功，返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"userID":  userID,
	})
}
