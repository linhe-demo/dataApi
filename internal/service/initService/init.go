package initService

import (
	"dataApi/app"
	"dataApi/internal/model/initModel"
	f "dataApi/pkg/festival"
	"dataApi/pkg/lunar"
	"dataApi/pkg/tools"
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
	Year          string `json:"year"`
	Month         string `json:"month"`
	NickName      string `json:"nick_name"`
	CoverImg      string `json:"cover_img"`
	FireWorkWords string `json:"fire_work_words"`
	BlessChinese  string `json:"bless_chinese"`
	BlessEnglish  string `json:"bless_english"`
	BtnWords      string `json:"btn_words"`
	BlessFestival string `json:"bless_festival"`
	FestivalName  string `json:"festival_name"`
	Flower        bool   `json:"flower"`
}

func GetInitData(ctx *gin.Context) (out ConfigData, err error) {
	var festivalConfig []initModel.Res
	mouldId := app.GetMouldID(ctx)
	out.NickName = app.GetNickName(ctx)
	//检测模板是否合法
	_, err = initModel.GetMouldInfo(mouldId)
	if err != nil {
		return out, errors.Wrap(err, "get mould config fail")
	}
	//获取模板固定配置
	mouldData, err := initModel.GetMouldData(mouldId)
	date := time.Now().Format(app.DateLayout)
	//获取今天的农历日期
	lunarDate := getLunarCalendar(date)
	out.Year = lunarDate.Year
	out.Month = lunarDate.Month
	//获取今天节日信息
	festival := getFestivalInfo(date)
	if len(festival) > app.DefaultInt { //获取节日配置
		festivalConfig, err = initModel.GetFestivalData(mouldId, festival[0])
		if festival[0] == "女王" {
			out.Flower = true
		}
		out.FestivalName = tools.InsertStringSpecialCharacter(festival[0], "|")
	} else {
		//获取默认烟花语
		firework, err := initModel.GetDefaultWordsInfo()
		if err != nil {
			return out, errors.Wrap(err, "get Default Words fail")
		}
		out.FireWorkWords = firework.Value
	}
	//检查节日配置是否存在
	if len(festivalConfig) == 0 {
		dealGreetingCardData(&out, mouldData)
	} else {
		dealGreetingCardData(&out, festivalConfig)
	}
	return out, nil
}

func dealGreetingCardData(out *ConfigData, data []initModel.Res) {
	for _, v := range data {
		switch v.Type {
		case 1:
			out.CoverImg = v.Value
		case 2:
			out.FireWorkWords = tools.InsertStringSpecialCharacter(v.Value, "|")
		case 3:
			out.BlessChinese = v.Value
		case 4:
			out.BlessEnglish = v.Value
		case 5:
			out.BtnWords = v.Value
		case 6:
			out.BlessFestival = v.Value
		case 7:
			out.FestivalName = v.Value
		}
	}
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
