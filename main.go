package main

import (
	"gin_jwt/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 路由
	router.Router(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
