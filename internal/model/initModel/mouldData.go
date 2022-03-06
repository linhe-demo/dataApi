package initModel

import (
	"dataApi/app"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// MouldData 模板信息表
type MouldData struct {
	ID         int       `gorm:"primary_key;column:id;type:int(10);not null"` // 记录id
	MouldID    int       `gorm:"index;column:mould_id;type:int(10);not null"` // 模板id
	Type       int16     `gorm:"column:type;type:smallint(6)"`                // 类型 1=封面图，3=中文祝福语，4=英文祝福语言，5=按钮文字，6=节日祝福语，7=默认节日名
	Value      string    `gorm:"column:value;type:varchar(200)"`              // 内容
	CreateDate time.Time `gorm:"column:create_date;type:datetime;not null"`   // 创建时间
	UpdateDate time.Time `gorm:"column:update_date;type:datetime;not null"`   // 更新时间
}

var MouldDataNotFound = errors.New("mould data not found")

func GetMouldData(mouldId int64) (out []Res, err error) {
	tx := app.MysqlClient.Model(MouldData{}).Where("mould_id = ?", int(mouldId)).Find(&out)
	if tx.Error == gorm.ErrRecordNotFound {
		return out, MouldDataNotFound
	} else if tx.Error != nil {
		return out, errors.Wrapf(err, "get mould data fail mouldId %d ", mouldId)
	}
	return out, nil
}
