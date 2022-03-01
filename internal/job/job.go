package job

import (
	"dataApi/conf"
	"dataApi/logs"
	"github.com/robfig/cron/v3"
)

type printfLogger struct {
}

func (pl printfLogger) Info(msg string, keysAndValues ...interface{}) {
	logs.Logger.Infow(msg, keysAndValues)
}

func (pl printfLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	logs.Logger.Errorw(msg, keysAndValues)
}
func JobRun() {
	if conf.AppConfig.Server.JobMaster != false {
		return
	}
	p := printfLogger{}
	//cronJob
	c := cron.New(cron.WithChain(
		cron.Recover(p), // or use cron.DefaultLogger
	))
	//_, _ = c.AddFunc("0 * * * *", xcWechatApi.RefreshXCToken)
	c.Start()
}
