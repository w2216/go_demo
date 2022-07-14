package router

import (
	"gin_jwt/controller"
	"gin_jwt/middelware"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	// jwt
	jwt := r.Group("/jwt")
	{
		//用户登录
		jwt.POST("/login", controller.LoginController)
		// 获取token
		jwt.GET("/index", controller.JwtController)
		//使用中间件
		jwt.Use(middelware.JWTAuth())
		//获取列表数据
		jwt.GET("/list", controller.UserListController)
	}

}
