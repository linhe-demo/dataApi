package controller

//common
const (
	CodeSuccess             = 1
	CodeMessage             = "success"
	CodeServerError         = 2
	CodeServerMessage       = "服务器错误"
	CodeParamFail           = 3
	CodeParamFailMessage    = "param fail"
	CodeParamIllegalCode    = 4
	CodeParamIllegalMessage = "param illegal"

	CodeParamIpFail        = 10001
	CodeParamIpFailMessage = "获取IP地址失败"

	// CodeRegisterFail 注册相关
	CodeRegisterFail        = 20001
	CodeRegisterFailMessage = "注册失败"
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
