package main

import (
	"dataApi/app"
	"dataApi/conf"
	"dataApi/internal/job"
	"dataApi/logs"
	appRouter "dataApi/router"
	"dataApi/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof"
	"time"
)

//程序入口
func main() {
	app.Init()

	if conf.AppConfig.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	//监控
	go appRouter.MetricsSrv()
	rout := gin.New()

	rout.Use(gin.Logger())
	rout.Use(gin.Recovery())
	//开启跨域
	rout.Use(middleware.Cors())

	//接入gin日志到zap中
	rout.Use(logs.Ginzap(logs.Logger, time.RFC3339, true))
	rout.Use(logs.RecoveryWithZap(logs.Logger, true))
	rout.Use(middleware.Metrics())

	//job 定时脚本
	go job.JobRun()

	//业务
	appRouter.RoutersRegister(rout)

	s := &http.Server{
		Addr:           conf.AppConfig.Server.HttpListen,
		Handler:        rout,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	logs.Logger.Errorw("app httpServer start err", "err", err)
}
