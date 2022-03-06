package router

import (
	"dataApi/api/controller/initData"
	"dataApi/api/controller/logs"
	"dataApi/api/controller/user"
	"dataApi/router/middleware"
	"github.com/gin-gonic/gin"
)

func RoutersRegister(r *gin.Engine) {
	//r.GET("/log/save", logs.Log)  //测试接口
	r.POST("/user/login", middleware.Chcek(), user.Login)   //用户登录
	r.POST("/user/register", user.RegisterAccount)          //用户注册
	r.POST("/log/save", middleware.Chcek(), logs.Log)       //记录访问日志
	r.POST("/init/info", middleware.Chcek(), initData.Info) //获取初始化数据
}
