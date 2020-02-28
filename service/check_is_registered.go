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

	count := 0
	//在搜索数据库，判断是否存在该用户
	if model.DB.Model(&model.UserInfo{}).Where("uid = ? and token = ?", service.Uid, service.Token).Count(&count); count == 0 {
		return serializer.BuildIsRegisteredResponse(0)
	}

	return serializer.BuildIsRegisteredResponse(1)
}
