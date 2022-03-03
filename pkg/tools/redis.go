package tools

import (
	"context"
	"dataApi/app"
	"dataApi/logs"
	"strconv"
)

//向redis中存入数据
func WriteInfoToUser(info []byte, roomid int64) {
	_, err := app.RedisClient.Publish(context.TODO(), strconv.FormatInt(roomid, 10), info).Result()
	if err != nil {
		logs.Logger.Errorw("WriteInfoToUser write info failed", "err", err, roomid, info)
	}
}
