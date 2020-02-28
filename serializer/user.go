package serializer

import "ncov_go/model"

// Usertoken 用户token器
type Status struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
}

// User 用户序列化器
type WeixinUser struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	PhoneNum string `json:"phone_num"`
	UserId   string `json:"userid"`
}

// Corp 表单模板序列化器
type Corp struct {
	Corpname     string `json:"corpname"`
	TypeCorpname string `json:"type_corpname"`
	TypeUsername string `json:"type_username"`
}

// isRegistered 用户序列化器
type IsRegistered struct {
	IsRegistered int `json:"is_registered"`
}

type CheckUser struct {
	IsExist int `json:"is_exist"`
}

//// BuildCorp 序列化status
//func BuildUserCheck(corp model.Corp) CheckUser {
//	return CheckUser{
//		IsExist:1,
//	}
//}

// BuildCorp 序列化status
func BuildUserInfo(user model.WeiXinUser) WeixinUser {
	return WeixinUser{
		Uid:      user.Uid,
		Name:     user.Name,
		PhoneNum: user.PhoneNum,
		UserId:   user.UserId,
	}
}

// BuildCorp 序列化status
func BuildCorp(corp model.Corp) Corp {
	return Corp{
		Corpname:     corp.Corpid,
		TypeCorpname: corp.TemplateCode,
		TypeUsername: corp.TemplateCode,
	}
}

// BuildStatus 序列化status
func BuildStatus(info model.Code) Status {
	return Status{
		Uid:   info.Uid,
		Token: info.Token,
	}
}

// BuildstatusResponse 序列化status响应
func BuildStatusResponse(info model.Code) Response {
	return Response{
		Data: BuildStatus(info),
	}
}

// BuildCorpResponse 序列化status响应
func BuildCorpResponse(corp model.Corp) Response {
	return Response{
		Data: BuildCorp(corp),
	}
}

// BuildstatusResponse 序列化status响应
func BuildIsRegisteredResponse(x int) Response {
	return Response{
		Data: IsRegistered{IsRegistered: x},
	}
}

// BuildstatusResponse 序列化status响应
func BuildUserCheckResponse(x int) Response {
	return Response{
		Data: CheckUser{IsExist: x},
	}
}

// BuildstatusResponse 序列化status响应
func BuildUserInfoResponse(user model.WeiXinUser) Response {
	return Response{
		Data: BuildUserInfo(user),
	}
}
