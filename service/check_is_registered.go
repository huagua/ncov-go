package service

import (
	"github.com/gin-gonic/gin"
	"ncov_go/model"
	"ncov_go/serializer"
)

// CheckIsRegisteredService 管理用户注册服务
type CheckIsRegisteredService struct {
	Code   string `form:"code" json:"code"`
	Corpid string `form:"corpid" json:"corpid"`
	Uid    string `form:"uid" json:"uid"`
	Token  string `form:"token" json:"token"`
}

// isRegistered 判断用户是否注册过
func (service *CheckIsRegisteredService) IsRegistered(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.Uid, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	return serializer.BuildIsRegisteredResponse(1)
}
