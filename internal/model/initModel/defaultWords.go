package initModel

import (
	"dataApi/app"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// DefaultWords [...]
type DefaultWords struct {
	ID    int    `gorm:"primary_key;column:id;type:int(10);not null"` // 记录id
	Value string `gorm:"column:value;type:varchar(255);not null"`     // 随机文字
}

var DefaultWordsNotFound = errors.New("default words not found")

func GetDefaultWordsInfo() (out DefaultWords, err error) {
	tx := app.MysqlClient.Model(DefaultWords{}).Order("RAND()").Limit(1).Find(&out)
	if tx.Error == gorm.ErrRecordNotFound {
		return out, DefaultWordsNotFound
	} else if tx.Error != nil {
		return out, errors.Wrapf(err, "get default words fail")
	}
	return out, nil
}
