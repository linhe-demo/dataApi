package user

import (
	"dataApi/api/controller"
	"dataApi/internal/service/userService"
	"dataApi/pkg/params"
	"github.com/gin-gonic/gin"
)

func RegisterAccount(ctx *gin.Context) {
	var data = userService.Param{}
	rsp := controller.MakeResponse()
	err := params.Unpack(ctx.Request, &data)
	if err != nil {
		rsp.Code = controller.CodeParamFail
		rsp.Msg = controller.CodeParamFailMessage
		rsp.Error = err
		ctx.JSON(200, rsp)
		return
	}
	if data.Mould <= 0 {
		rsp.Code = controller.CodeParamIllegalCode
		rsp.Msg = controller.CodeParamIllegalMessage
		ctx.JSON(200, rsp)
		return
	}
	res, err := userService.Register(data)
	if err != nil {
		rsp.Code = controller.CodeServerError
		rsp.Msg = controller.CodeServerMessage
		rsp.Error = err
		ctx.JSON(200, rsp)
		return
	}
	rsp.Code = controller.CodeSuccess
	rsp.Msg = controller.CodeMessage
	rsp.Data = res
	ctx.JSON(200, rsp)
	return
}
