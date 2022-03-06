package initModel

import (
	"dataApi/app"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// MouldFestival 模板节日信息表
type MouldFestival struct {
	ID         int       `gorm:"primary_key;column:id;type:int(10);not null"` // 记录id
	MouldID    int       `gorm:"index;column:mould_id;type:int(10);not null"` // 模板id
	Type       int16     `gorm:"column:type;type:smallint(6)"`                // 类型 1=封面图，2=烟花展示语，3=中文祝福语，4=英文祝福语言，5=按钮文字，6=节日祝福语，7=节日名
	Value      string    `gorm:"column:value;type:varchar(200)"`              // 内容
	CreateDate time.Time `gorm:"column:create_date;type:datetime;not null"`   // 创建时间
	UpdateDate time.Time `gorm:"column:update_date;type:datetime;not null"`   // 更新时间
}

type Res struct {
	Type  int16
	Value string
}

var MouldFestivalNotFound = errors.New("mould festival not found")

func GetFestivalData(mouldId int64, festivalName string) (out []Res, err error) {
	tx := app.MysqlClient.Model(MouldFestival{}).Where("MouldID = ? AND Value = ?", int(mouldId), festivalName).Find(&out)
	if tx.Error == gorm.ErrRecordNotFound {
		return out, MouldFestivalNotFound
	} else if tx.Error != nil {
		return out, errors.Wrapf(err, "get festival config fail mouldId %d festivalName %s", mouldId, festivalName)
	}
	return out, nil
}
