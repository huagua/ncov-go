package serializer

import "ncov_go/model"

// User 用户序列化器
type Status struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
}

// BuildStatus 序列化status
func BuildStatus(info model.UserInfo) Status {
	return Status{
		Uid:   info.Uid,
		Token: info.Token,
	}
}

// BuildstatusResponse 序列化status响应
func BuildStatusResponse(info model.UserInfo) Response {
	return Response{
		Data: BuildStatus(info),
	}
}

// BuildstatusResponse 序列化status响应
func BuildIsRegisteredResponse() Response {
	return Response{
		Data: true,
	}
}
