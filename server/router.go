package server

import (
	"github.com/gin-gonic/gin"
	"ncov_go/api"
)

func NewRouter() *gin.Engine {
	//创建带有默认中间件的路由
	//日志与恢复中间件
	router := gin.Default()

	// 路由
	v1 := router.Group("/api")
	{
		//获取uid和token
		v1.POST("/login/getcode", api.UserLogin)

		//检查用户是否注册
		v1.POST("/login/check_is_registered", api.UserIsReg)

		//检查用户是否存在
		v1.POST("/login/check_user", api.CheckUser)

		//微信用户注册
		v1.POST("/login/register", api.WeixinUsrRegister)

		//微信用户注册
		v1.POST("/info/getmyinfo", api.GetUserInfo)

		//获取公司模板
		v1.POST("/login/getcorpname", api.GetCorp)

		//保存每日上传信息
		v1.POST("/report/save", api.SaveInfo)

		//获取每日上传信息
		v1.POST("/report/getlastdata", api.GetInfo)
	}

	return router
}
