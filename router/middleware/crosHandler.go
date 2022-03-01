package middleware

import (
	"github.com/gin-gonic/gin"
)

// 开启跨域函数
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		defer func() {
			if err := recover(); err != nil {
				//logs.Logger.Errorw("Panic info is: %v", "err", err)
				//logs.Logger.Errorw("Panic info is: %s", "err", debug.Stack())
			}
		}()
		c.Next()
	}
}