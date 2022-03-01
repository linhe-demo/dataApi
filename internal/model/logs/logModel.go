package logs

import (
	"dataApi/app"
	"github.com/pkg/errors"
	"time"
)

// VisitLog [...]
type VisitLog struct {
	ID   int       `gorm:"primary_key;column:id;type:int(10);not null"` // 记录id
	IP   string    `gorm:"column:ip;type:varchar(255);not null"`        // ip地址
	Date time.Time `gorm:"column:date;type:datetime;not null"`          // 时间
}

func InsertLog(ip string) (out bool, err error) {
	info := VisitLog{
		IP:   ip,
		Date: time.Now(),
	}
	tx := app.MysqlClient.Model(VisitLog{}).Create(&info)

	if tx.Error != nil {
		return out, errors.Wrapf(tx.Error, "save log info failed ip %s date %s", ip, time.Now().Format(app.Datelayouttime))
	}
	return true, nil
}