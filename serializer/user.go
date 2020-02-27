package serializer

import "ncov_go/model"

// User 用户序列化器
type Status struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
}

// Corp 用户序列化器
type Corp struct {
	Corpid       string `json:"corpid"`
	TemplateCode string `json:"template_code"`
}

// BuildCorp 序列化status
func BuildCorp(corp model.Corp) Corp {
	return Corp{
		Corpid:       corp.Corpid,
		TemplateCode: corp.TemplateCode,
	}
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

// BuildCorpResponse 序列化status响应
func BuildCorpResponse(corp model.Corp) Response {
	return Response{
		Data: BuildCorp(corp),
	}
}

// BuildstatusResponse 序列化status响应
func BuildIsRegisteredResponse() Response {
	return Response{
		Data: true,
	}
}
