package api

import (
	"github.com/gin-gonic/gin"
	"ncov_go/service"
)

// UserLogin 用户登录接口，获取openid，token
func UserLogin(c *gin.Context) {
	var service service.UserOpenIdService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetCode(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserIsReg 判断用户是否存在
func UserIsReg(c *gin.Context) {
	var service service.CheckIsRegisteredService
	if err := c.ShouldBind(&service); err == nil {
		res := service.IsRegistered(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// SaveInfo 用户上传信息接口
func SaveInfo(c *gin.Context) {
	var service service.SaveDailyInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SaveDailyInfo(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetInfo 用户上传信息接口
func GetInfo(c *gin.Context) {
	var service service.GetInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetLastData(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetCorp 用户上传信息接口
func GetCorp(c *gin.Context) {
	var service service.GetCorpService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetCorp(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
