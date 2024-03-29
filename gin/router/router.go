package router

import (
	"gin/controller"
	"gin/controller/alipay"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.GET("", controller.IndexController)

	// jwt
	jwt := r.Group("/jwt")
	{
		//用户登录
		jwt.POST("/login", controller.LoginController)
		// 获取token
		jwt.GET("/index", controller.JwtController)
		//使用中间件
		jwt.Use(middleware.JWTAuth())
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

	// web
	web := r.Group("/web")
	{
		//注册
		web.POST("/register", controller.RegisterController)

		//登录
		web.POST("/login", controller.LoginWebController)

	}

	// chan
	inCh := r.Group("/chan")
	{
		//流水线模型1
		inCh.POST("/chan1", controller.ChanController)
		//流水线模型2
		inCh.POST("/chan2", controller.Chan2Controller)
	}

	// ali
	ali := r.Group("/alipay")
	{
		// 创建订单
		ali.POST("/create", alipay.CreateController)
		// 下单
		ali.POST("/pay", alipay.PayController)
		// 退款
		ali.POST("/refund", alipay.RefundController)
		// 回调通知
		ali.POST("/notify", alipay.NotifyController)
	}

}
