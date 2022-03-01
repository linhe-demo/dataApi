package router

import (
	"dataApi/api/controller/logs"
	"github.com/gin-gonic/gin"
)

func RoutersRegister(r *gin.Engine) {
	//r.GET("/log/save", logs.Log)  //测试接口
	r.POST("/log/save", logs.Log) //测试接口
}
