package initModel

import (
	"dataApi/app"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// MouldConfig 模板表
type MouldConfig struct {
	ID         int       `gorm:"primary_key;column:id;type:int(10);not null"` // 模板id
	Name       string    `gorm:"column:name;type:varchar(100)"`               // 名称
	Descrption string    `gorm:"column:descrption;type:varchar(500)"`         // 描述
	CreateDate time.Time `gorm:"column:create_date;type:datetime;not null"`   // 创建时间
	UpdateDate time.Time `gorm:"column:update_date;type:datetime;not null"`   // 更新时间
}

var MouldConfigNotFound = errors.New("mould config not found")

func GetMouldInfo(mouldId int64) (out MouldConfig, err error) {
	tx := app.MysqlClient.Where("ID = ?", int(mouldId)).Find(&out)
	if tx.Error == gorm.ErrRecordNotFound {
		return out, MouldConfigNotFound
	} else if tx.Error != nil {
		return out, errors.Wrapf(err, "get mould info fail mouldId %d", mouldId)
	}
	return out, nil
}
