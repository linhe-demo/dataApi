package middleware

import (
	"context"
	"dataApi/app"
	"dataApi/pkg/eado"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type CheckInfo struct {
	Mould    int64  `json:"mould"`
	Time     int64  `json:"time"`
	UserID   int64  `json:"user_id"`
	RandNum  int64  `json:"rand_num"`
	NickName string `json:"nick_name"`
}

// 开启token校验
func Chcek() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		index := strings.Index(url, "?")
		if index != -1 {
			rs := []rune(url)
			url = string(rs[0:index])
		}
		token := c.Request.Header["Access-Token"]
		if len(token) > 0 && len(token[0]) > 0 {
			//解析数据
			clientTimeDecode, err := base64.StdEncoding.DecodeString(token[0])
			if err != nil {
				c.AbortWithStatus(401)
				return
			}
			redisKeyXor := eado.Xor(clientTimeDecode, []byte(app.LoginKey))
			redisKey := string(redisKeyXor)
			tmpReadisKey := strings.Split(redisKey, "|||")
			redisKey = tmpReadisKey[1]
			tmpRandNum, _ := strconv.ParseInt(tmpReadisKey[2], 10, 64)

			//获取redis中数据
			cmd := app.RedisClient.Get(context.TODO(), redisKey)
			if err := cmd.Err(); err != nil {
				c.AbortWithStatus(401)
				return
			} else {
				jsonStr, _ := cmd.Result()
				var mapResult CheckInfo
				err := json.Unmarshal([]byte(jsonStr), &mapResult)
				if err == nil {
					//判断用户登录是否超过两天
					randNum := mapResult.RandNum
					timeNow := time.Now().Unix()
					beginTime := mapResult.Time
					tmpUserid := mapResult.UserID
					tmpMould := mapResult.Mould
					tmpNickName := mapResult.NickName
					if randNum != tmpRandNum {
						c.AbortWithStatus(401)
						return
					}
					if timeNow-beginTime > 86400*7 { //登录token已过期
						c.AbortWithStatus(401)
						return
					} else {
						if tmpUserid == 0 {
							c.AbortWithStatus(401)
							return
						}
						c.Set("userinfo", app.UserInfo{Uid: tmpUserid, Mould: tmpMould, NickName: tmpNickName})
					}
				} else {
					c.AbortWithStatus(401)
					return
				}
			}
		} else {
			c.AbortWithStatus(401)
			return
		}
	}
}
