package userModel

import (
	"dataApi/app"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// UserInfo 用户信息表
type UserInfo struct {
	ID         int       `gorm:"primary_key;column:id;type:int(10);not null"`       // 记录id
	MouldID    int       `gorm:"index;column:mould_id;type:int(10);not null"`       // 模板id
	NickName   string    `gorm:"index;column:nick_name;type:varchar(100);not null"` // 用户昵称
	CreateDate time.Time `gorm:"column:create_date;type:datetime;not null"`         // 创建时间
	LoginDate  time.Time `gorm:"column:login_date;type:datetime;not null"`          // 登录时间
}

var UserInfoNotFound = errors.New("user info not found")

func GetUserInfo(uid int64, mouldId int64) (out UserInfo, err error) {
	tx := app.MysqlClient.Where("id = ? AND mould_id = ?", int(uid), int(mouldId)).First(&out)
	if tx.Error == gorm.ErrRecordNotFound {
		return out, UserInfoNotFound
	} else if tx.Error != nil {
		return out, errors.Wrapf(err, "get user info fail uid %d mouldId %d", uid, mouldId)
	}
	return out, nil
}

func UpdateUserLoginInfo(uid int64, mouldId int64) {
	app.MysqlClient.Model(UserInfo{}).Where("id = ? AND mould_id = ?", int(uid), int(mouldId)).Update("login_date", time.Now().Format(app.DatelayoutTime))
}

func CreateUserInfo(mouldId int64, nickName string) (uid int, err error) {
	info := UserInfo{
		MouldID:    int(mouldId),
		NickName:   nickName,
		CreateDate: time.Now(),
		LoginDate:  time.Now(),
	}
	tx := app.MysqlClient.Model(UserInfo{}).Create(&info)

	if tx.Error != nil {
		return uid, errors.Wrapf(tx.Error, "create user account failed nickname %s mouldId %d", nickName, mouldId)
	}
	return info.ID, nil
}

func GetUserInfoByNickName(nickName string) (out UserInfo, err error) {
	tx := app.MysqlClient.Where("nick_name = ?", nickName).First(&out)
	if tx.Error == gorm.ErrRecordNotFound {
		return out, UserInfoNotFound
	} else if tx.Error != nil {
		return out, errors.Wrapf(err, "get user info fail mickname %s", nickName)
	}
	return out, nil
}
