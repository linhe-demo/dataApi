package logs

import (
	"dataApi/api/controller"
	"dataApi/internal/service/logService"
	"github.com/gin-gonic/gin"
)

func Log(ctx *gin.Context) {
	rsp := controller.MakeResponse()
	ip := ctx.ClientIP()
	if len(ip) == 0 {
		rsp.Code = controller.CodeParamIpFail
		rsp.Msg = controller.CodeParamIpFailMessage
		ctx.JSON(200, rsp)
		return
	}

	_, err := logService.SaveLog(ctx, ip)
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
