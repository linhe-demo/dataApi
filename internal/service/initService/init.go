package initService

import (
	"dataApi/app"
	"dataApi/internal/model/initModel"
	f "dataApi/pkg/festival"
	"dataApi/pkg/lunar"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type Param struct {
	Mould int64 `json:"mould"`
}

type LunarDate struct {
	Year  string
	Month string
}

type ConfigData struct {
}

func GetInitData(ctx *gin.Context) (out ConfigData, err error) {
	var (
		festivalConfig []initModel.MouldFestival
	)
	uid := app.GetUserId(ctx)
	mouldId := app.GetMouldID(ctx)
	//检测模板是否合法
	config, err := initModel.GetMouldInfo(mouldId)
	if err != nil {
		return out, errors.Wrap(err, "get mould config fail")
	}
	//获取模板固定配置
	date := time.Now().Format(app.DateLayout)
	//获取今天的农历日期
	lunarDate := getLunarCalendar(date)
	//获取今天节日信息
	festival := getFestivalInfo(date)
	if len(festival) > app.DefaultInt { //获取节日配置
		for _, v := range festival {
			tmp, err := initModel.GetFestivalData(mouldId, v)
			if err == nil && tmp.ID > 0 {
				festivalConfig = append(festivalConfig, tmp)
			}
		}
	}
	return out, nil
}

func getLunarCalendar(date string) (out LunarDate) {
	lunarDate := lunar.SolarToChineseLuanr(date)
	tmpSlice := strings.Split(lunarDate, "|")
	out.Year = tmpSlice[0]
	out.Month = tmpSlice[1]
	return out
}

func getFestivalInfo(date string) (out []string) {
	festival := f.NewFestival("../pkg/festival/festival.json")
	out = festival.GetFestivals(date)
	return out
}
