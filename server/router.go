package server

import (
	"github.com/gin-gonic/gin"
	"ncov_go/api"
)

func NewRouter() *gin.Engine {
	//创建带有默认中间件的路由
	//日志与恢复中间件
	router := gin.Default()

	// 路由 用户登录信息接口
	v1 := router.Group("/login")
	{
		//获取uid和token
		v1.POST("/getcode", api.UserLogin)

		//检查用户是否注册
		v1.POST("/check_is_registered", api.UserIsReg)

		//检查用户是否存在
		v1.POST("/check_user", api.CheckUser)

		//微信用户注册
		v1.POST("/register", api.WeixinUsrRegister)

		//获取公司模板
		v1.POST("/getcorpname", api.GetCorp)

		//绑定用户
		v1.POST("/bind", api.UserBind)

		//用户解绑
		v1.POST("/unbind", api.UserUnBind)
	}

	// 路由，用户上传信息接口
	v2 := router.Group("/report")
	{
		//保存每日上传信息
		v2.POST("/save", api.SaveInfo)

		//获取上次上传信息
		v2.POST("/getlastdata", api.GetInfo)
	}

	// 路由， 获取用户基本信息
	v3 := router.Group("/info")
	{
		//获取用户信息
		v3.POST("/getmyinfo", api.GetUserInfo)
	}

	return router
}
