package logModel

import (
	"dataApi/app"
	"github.com/pkg/errors"
	"time"
)

// VisitLog 用户访问日志表
type VisitLog struct {
	ID      int       `gorm:"primary_key;column:id;type:int(10);not null"` // 记录id
	MouldID int       `gorm:"index;column:mould_id;type:int(10);not null"` // 模板id
	UId     int       `gorm:"index;column:uid;type:int(10);not null"`      // 用户id
	IP      string    `gorm:"column:ip;type:varchar(255);not null"`        // ip地址
	Date    time.Time `gorm:"column:date;type:datetime;not null"`          // 时间
}

func InsertLog(uid int64, mouldId int64, ip string) (out bool, err error) {
	info := VisitLog{
		MouldID: int(mouldId),
		UId:     int(uid),
		IP:      ip,
		Date:    time.Now(),
	}
	tx := app.MysqlClient.Model(VisitLog{}).Create(&info)

	if tx.Error != nil {
		return out, errors.Wrapf(tx.Error, "save log info failed ip %s date %s", ip, time.Now().Format(app.DatelayoutTime))
	}
	return true, nil
}
