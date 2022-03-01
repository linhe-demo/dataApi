package router

import (
	"dataApi/conf"
	"dataApi/logs"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"net/http/pprof"
	"time"
)

const (
	// DefaultPrefix url prefix of pprof
	DefaultPrefix = "/debug/pprof"
)

func getPrefix(prefixOptions ...string) string {
	prefix := DefaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}
	return prefix
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func registerPprof(r *gin.Engine, prefixOptions ...string) {
	prefix := getPrefix(prefixOptions...)

	prefixRouter := r.Group(prefix)
	{
		prefixRouter.GET("/", pprofHandler(pprof.Index))
		prefixRouter.GET("/cmdline", pprofHandler(pprof.Cmdline))
		prefixRouter.GET("/profile", pprofHandler(pprof.Profile))
		prefixRouter.POST("/symbol", pprofHandler(pprof.Symbol))
		prefixRouter.GET("/symbol", pprofHandler(pprof.Symbol))
		prefixRouter.GET("/trace", pprofHandler(pprof.Trace))
		prefixRouter.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		prefixRouter.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		prefixRouter.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		prefixRouter.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		prefixRouter.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		prefixRouter.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	}
}

func MetricsSrv() {
	defer func() {
		if err := recover(); err != nil {
			logs.Logger.Errorw("metrics Srv panic", "err", err)
		}
	}()

	router := gin.New()
	router.Use(logs.RecoveryWithZap(logs.Logger, true))

	//prom监控
	router.GET("metrics", gin.WrapH(promhttp.Handler()))
	//if conf.AppConfig.Server.Debug {
	registerPprof(router)
	//}

	s := &http.Server{
		Addr:           conf.AppConfig.Server.PrivateHttpListen,
		Handler:        router,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	logs.Logger.Infow("app metric httpServer start err", "err", err)
}
