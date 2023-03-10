package initialize

import (
	"fmt"
	"gin-api-server/conf"
	"gin-api-server/tools"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var Log *zap.Logger

func InitLogger() {
	fmt.Println(conf.Zap.DebugFileName)
	writeSyncerDebug := getLogWriter(conf.Zap.DebugFileName, conf.Zap.MaxSize, conf.Zap.MaxBackups, conf.Zap.MaxAge)
	writeSyncerInfo := getLogWriter(conf.Zap.InfoFileName, conf.Zap.MaxSize, conf.Zap.MaxBackups, conf.Zap.MaxAge)
	writeSyncerWarn := getLogWriter(conf.Zap.WarnFileName, conf.Zap.MaxSize, conf.Zap.MaxBackups, conf.Zap.MaxAge)
	encoder := getEncoder()
	//文件输出
	debugCore := zapcore.NewCore(encoder, writeSyncerDebug, zapcore.DebugLevel)
	infoCore := zapcore.NewCore(encoder, writeSyncerInfo, zapcore.InfoLevel)
	warnCore := zapcore.NewCore(encoder, writeSyncerWarn, zapcore.WarnLevel)
	//标准输出
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	std := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	core := zapcore.NewTee(debugCore, infoCore, warnCore, std)
	Log = zap.New(core, zap.AddCaller())
	//替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(Log)
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logid := strings.Join([]string{"520", strconv.FormatInt(time.Now().UnixNano(), 10)}, "")
		c.Set("logid", logid)

		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)

		body := make([]byte, c.Request.ContentLength)
		_, err := c.Request.Body.Read(body)
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		Log.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("ip", c.ClientIP()),
			zap.String("query", query),
			zap.String("body", string(body)),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
			zap.String("logid", logid),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(logid string) {
			if err := recover(); err != nil {

				tools.SendEmailWarn(conf.Server.ServerName, logid)
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					Log.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("logid", logid),
					)
					c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					Log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
						zap.String("logid", logid),
					)
				} else {
					Log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("logid", logid),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}(c.GetString("logid"))
		c.Next()
	}
}
