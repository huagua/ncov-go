package service

import (
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
	"ncov_go/model"
	"ncov_go/serializer"
	"os"
)

// UserOpenIdService 获取用户token服务
type UserOpenIdService struct {
	Code string `form:"code" json:"code"`
}

// gecode 用户登录函数，获取openidhesessionkey，作为之后操作的验证
func (service *UserOpenIdService) GetCode(c *gin.Context) serializer.Response {
	res, err := weapp.Login(os.Getenv("APP_ID"), os.Getenv("APP_SECREAT"), service.Code)

	if err != nil {
		//处理错误
		return serializer.ParamErr("获取openid失败", err)
	}

	if err := res.GetResponseError(); err != nil {
		//处理小程序传送的错误信息
		return serializer.ParamErr("小程序报错", err)
	}

	info := model.UserInfo{
		Uid:    res.OpenID,
		Token:  res.SessionKey,
		Corpid: "100000001",
		Code:   service.Code,
	}

	// 记录用户token信息
	if err := model.DB.Create(&info).Error; err != nil {
		return serializer.ParamErr("记录失败", err)
	}

	return serializer.BuildStatusResponse(info)
}
