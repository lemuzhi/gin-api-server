package middleware

import (
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//s := "panic recover err: %v"
				//global.Logger
				//TODO 处理逻辑，日志写入，告警邮件发送等
				c.Abort()
			}
		}()
		c.Next()
	}
}
