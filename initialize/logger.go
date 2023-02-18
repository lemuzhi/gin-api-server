package initialize

import (
	"fmt"
	"gin-project-template/global"
	"gin-project-template/pkg/format"
	"gin-project-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"strings"
)

func InitLogger() {
	//设置日志的输出格式 ：本地时区中的日期 完整文件名和行号
	//log.SetFlags(log.Ltime | log.Ldate | log.Llongfile)
	filePath := viper.GetString("log.file_path")
	name := strings.ReplaceAll(format.GetTime(format.TimeOnly, ""), ":", "")
	fileName := filePath + name + ".log"
	fmt.Println(fileName)
	logFile, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
	}
	//lumberjack日志滚动记录器
	//lumberjack是一个日志滚动记录器。可以把日志文件根据大小、日期等分割。一般情况下，
	//lumberjack配合其他日志库，实现日志的滚动(rolling)记录。
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,                       //日志文件的位置
		MaxSize:    viper.GetInt("log.max_size"),   //切割之前日志文件得大小（单位：MB）
		MaxBackups: viper.GetInt("log.ma_backups"), //保留旧文件得最大个数
		MaxAge:     viper.GetInt("log.max_age"),    //保留旧文件得最大天数
		Compress:   true,                           //是否压缩旧文件
	}
	global.Log = logger.NewLogger(lumberJackLogger, "", log.LstdFlags).WithCaller(2)
	log.SetOutput(lumberJackLogger)
	//同时写文件和屏幕
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}
