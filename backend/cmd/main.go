package main

import (
	"backend/api/login"
	run_ocr "backend/api/run-ocr"
	"backend/api/tr-run"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	port        int = 8089
	altPort     int = 8000
	frontendURL     = "http://localhost:8080" // 假设前端服务运行在3000端口上
)

func main() {

	r := gin.Default()

	// 设置跨域中间件colons
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{frontendURL}, // 允许指定的前端地址跨域访问
		AllowMethods: []string{"POST"},      // 允许POST方法跨域访问
		//AllowHeaders: []string{"Origin"},    // 允许指定的请求头跨域访问
		AllowHeaders: []string{"Content-Type"}, // 允许Content-Type请求头跨域访问
	}))

	// 设置API路由组
	api := r.Group("/api")
	{
		api.POST("/tr-run", tr_run.TrRun)  // 使用 tr_run 包中的 TrRun 函数处理 POST 请求
		api.POST("/sign-in", login.SignIn) //登录
		api.POST("/sign-up", login.SignUp) //注册
	}

	// 启动主服务
	go func() {
		//fmt.Printf("Server is running: http://%s:%d\n", utils.HostIP(), port)
		r.Run("localhost" + fmt.Sprintf(":%d", port)) // 启动主服务
	}()

	// 启动额外服务
	go func() {
		r2 := gin.Default()
		api2 := r2.Group("/api")
		{
			// 这里添加额外服务的路由
			api2.GET("/run-ocr", run_ocr.RunPythonScript)
		}
		//fmt.Printf("Server is running: http://%s:%d\n", utils.HostIP(), altPort)
		r2.Run("localhost" + fmt.Sprintf(":%d", altPort)) // 启动额外服务
	}()

	// 阻塞主 goroutine，保持服务运行
	select {}
}
