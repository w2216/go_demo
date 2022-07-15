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

	// gorm
	gorm := r.Group("/gorm")
	{
		//创建表
		gorm.POST("/create", controller.CreateController)

		//插入数据
		gorm.POST("/insert", controller.InsertController)

		//更新数据
		gorm.POST("/update", controller.UpdateController)

		//查询数据
		gorm.POST("/select", controller.SelectController)

		//删除数据
		gorm.POST("/delete", controller.DeleteController)

	}

   // redis
	redis := r.Group("/redis")
	{
		//设置值
		redis.POST("/set", controller.RdbSetController)

      //获取值
		redis.POST("/get", controller.RdbGetController)

      //删除值
		redis.POST("/del", controller.RdbDelController)


	}

}
