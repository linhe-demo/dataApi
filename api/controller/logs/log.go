package logs

import (
	"dataApi/api/controller"
	"dataApi/internal/service/logService"
	"dataApi/pkg/params"
	"github.com/gin-gonic/gin"
)

func Log(ctx *gin.Context) {
	var data = logService.Param{}
	rsp := controller.MakeResponse()
	err := params.Unpack(ctx.Request, &data)
	if err != nil {
		rsp.Code = controller.CodeParamFail
		rsp.Msg = controller.CodeParamFailMessage
		rsp.Error = err
		ctx.JSON(200, rsp)
		return
	}
	data.Ip = ctx.ClientIP()
	if len(data.Ip) == 0 {
		rsp.Code = controller.CodeParamIpFail
		rsp.Msg = controller.CodeParamIpFailMessage
		ctx.JSON(200, rsp)
		return
	}

	_, err = logService.SaveLog(data)
	if err != nil {
		rsp.Code = controller.CodeServerError
		rsp.Msg = controller.CodeServerMessage
		rsp.Error = err
		ctx.JSON(200, rsp)
		return
	}
	rsp.Code = controller.CodeSuccess
	rsp.Msg = controller.CodeMessage
	ctx.JSON(200, rsp)
	return
}
