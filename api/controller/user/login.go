package user

import (
	"dataApi/api/controller"
	"dataApi/internal/service/userService"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	rsp := controller.MakeResponse()
	res, err := userService.CheckUserInfo(ctx)
	if err != nil {
		rsp.Code = controller.CodeServerError
		rsp.Msg = controller.CodeServerMessage
		rsp.Error = err
		ctx.JSON(200, rsp)
		return
	}
	rsp.Code = controller.CodeSuccess
	rsp.Msg = controller.CodeMessage
	rsp.Data = res.Message
	ctx.JSON(200, rsp)
	return
}
