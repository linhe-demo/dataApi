package logService

import (
	"dataApi/app"
	"dataApi/internal/model/logModel"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func SaveLog(ctx *gin.Context, ip string) (out bool, err error) {
	uid := app.GetUserId(ctx)
	mouldId := app.GetMouldID(ctx)
	out, err = logModel.InsertLog(uid, mouldId, ip)
	if err != nil {
		return out, errors.Wrap(err, "logService save failed")
	}
	return true, nil
}
