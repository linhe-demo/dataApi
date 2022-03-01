package controller

//common
const (
	CodeSuccess          = 1
	CodeMessage          = "success"
	CodeServerError      = 2
	CodeServerMessage    = "服务器错误"
	CodeParamFail        = 3
	CodeParamFailMessage = "param fail"

	CodeParamIpFail        = 1001
	CodeParamIpFailMessage = "获取IP地址失败"
)

type HttpResponse struct {
	Code  int32       `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}

func MakeResponse() (rsp HttpResponse) {
	rsp.Code = CodeSuccess
	rsp.Msg = CodeMessage
	return rsp
}
