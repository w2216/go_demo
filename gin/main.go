package main

import (
	"gin/middleware"
	"gin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 记录日志
	r.Use(middleware.LoggerToFile())

	// 路由
	router.Router(r)
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}
