package user

import (
	"context"
	"dataApi/api/controller"
	"dataApi/app"
	"dataApi/internal/service/userService"
	"dataApi/pkg/eado"
	"dataApi/pkg/params"
	"dataApi/pkg/tools"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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
	if data.Mould <= app.DefaultInt || len(data.NickName) == app.DefaultInt {
		rsp.Code = controller.CodeParamIllegalCode
		rsp.Msg = controller.CodeParamIllegalMessage
		ctx.JSON(200, rsp)
		return
	}
	uid, err := userService.Register(data)
	if err != nil {
		rsp.Code = controller.CodeServerError
		rsp.Msg = controller.CodeServerMessage
		rsp.Error = err
		ctx.JSON(200, rsp)
		return
	}

	randNmu := tools.GenerateRandom()
	redisKey := app.UserRedisKey + strconv.Itoa(uid)
	tmpTimeStamp := time.Now().Unix()
	//加密数据
	secretData := map[string]interface{}{"mould": data.Mould, "time": tmpTimeStamp, "user_id": uid, "rand_num": randNmu, "nick_name": data.NickName}
	mjson, err := json.Marshal(secretData)
	if err == nil {
		app.RedisClient.Set(context.TODO(), redisKey, mjson, time.Second*86400*365)
		xor := eado.Xor([]byte(strconv.FormatInt(int64(tools.GenerateRandom()), 10)+"|||"+redisKey+"|||"+strconv.FormatInt(int64(randNmu), 10)+"|||"+strconv.FormatInt(tmpTimeStamp, 10)), []byte(app.LoginKey))
		token := base64.StdEncoding.EncodeToString(xor)
		rsp.Code = controller.CodeSuccess
		rsp.Msg = controller.CodeMessage
		rsp.Data = token
		ctx.JSON(200, rsp)
		return
	} else {
		rsp.Code = controller.CodeRegisterFail
		rsp.Msg = controller.CodeRegisterFailMessage
		ctx.JSON(200, rsp)
		return
	}
}
