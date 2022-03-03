package initData

import (
	"dataApi/api/controller"
	"dataApi/internal/service/initService"
	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	rsp := controller.MakeResponse()
	res, err := initService.GetInitData(ctx)
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
