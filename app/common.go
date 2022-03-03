package app

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func GetUserId(ctx *gin.Context) int64 {
	userinfo, exist := ctx.Get("userinfo")
	var userid int64
	if exist == true {
		userid = userinfo.(UserInfo).Uid
	}
	return userid
}

func GetMouldID(ctx *gin.Context) int64 {
	userinfo, exist := ctx.Get("userinfo")
	var mouldId int64
	if exist == true {
		mouldId = userinfo.(UserInfo).Mould
	}
	return mouldId
}

func Wait(num float32) {
	tmpTime := time.Duration(num * 1000000000)
	time.Sleep(tmpTime)
}

func BuildHttpUrl(baseUrl string, data map[string]string) string {
	if len(data) <= 0 {
		return baseUrl
	}
	baseUrl += "?"
	for k, v := range data {
		baseUrl += k + "=" + v + "&"
	}
	return strings.Trim(baseUrl, "&")
}

func Hash(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

//func QiniuToken() string {
//	putPolicy := storage.PutPolicy{
//		Scope:   conf.AppConfig.Qiniu.Bucket,
//		Expires: 7200, //2小时有效期
//	}
//	mac := qbox.NewMac(conf.AppConfig.Qiniu.Ak, conf.AppConfig.Qiniu.Sk)
//
//	return putPolicy.UploadToken(mac)
//}
