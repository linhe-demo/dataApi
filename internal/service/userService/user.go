package userService

import (
	"dataApi/app"
	"dataApi/internal/model/userModel"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Param struct {
	Mould    int64  `json:"mould"`
	NickName string `json:"nick_name"`
}

type BackInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CheckUserInfo(ctx *gin.Context) (out BackInfo, err error) {
	uid := app.GetUserId(ctx)
	mouldId := app.GetMouldID(ctx)
	if uid == app.DefaultInt {
		out.Code = app.DefaultInt
		out.Message = "用户id不存在"
		return out, nil
	}
	_, err = userModel.GetUserInfo(uid, mouldId)
	if err != nil {
		out.Code = app.DefaultInt
		out.Message = "用户id不存在"
		return out, errors.Wrap(err, "get user info fail")
	}
	//更新用户登录信息
	userModel.UpdateUserLoginInfo(uid, mouldId)
	out.Code = app.SuccessCode
	out.Message = "登录成功"
	return out, nil
}

func Register(info Param) (uid int, err error) {
	//检测昵称是否已存在
	userInfo, err := userModel.GetUserInfoByNickName(info.NickName)
	if err == userModel.UserInfoNotFound {
		uid, err = userModel.CreateUserInfo(info.Mould, info.NickName)
		if err != nil {
			return uid, errors.Wrap(err, "创建用户失败")
		}
	} else if err != nil {
		return uid, errors.Wrap(err, "网络错误，请重试")
	} else {
		uid = userInfo.ID
	}
	return uid, nil
}
