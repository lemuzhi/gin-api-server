package cmd

import (
	"context"
	"fmt"
	"gin-project-template/global"
	"gin-project-template/initialize"
	"gin-project-template/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

type ServerStartOptions struct {
	config string
}

func NewServerStartCmd(ctx context.Context, version string) *cobra.Command {
	opts := &ServerStartOptions{}

	cmd := &cobra.Command{
		Use:   "api",
		Short: "Start a api Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServerStart(ctx, opts, version)
		},
	}
	cmd.PersistentFlags().StringVarP(&opts.config, "config", "c", "./conf/config.toml", "Config file")
	return cmd
}

func RunServerStart(ctx context.Context, opts *ServerStartOptions, version string) error {
	//初始化配置
	initialize.InitConfig("./config/config.toml")

	//初始化日志
	initialize.InitLogger()

	//初始化mysql
	initialize.InitMysql()

	//初始化redis
	initialize.InitRedis()

	//初始化定时任务
	initialize.InitCron()

	//设置gin的启动模式
	gin.SetMode(viper.GetString("gin.mode"))

	//初始化路由
	handler := router.InitRouter()

	//自定义HTTP配置
	server := &http.Server{
		Addr:           global.Config.GetString("gin.addr"),
		Handler:        handler,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)

	fmt.Printf(`
		贵特100--服务端
		当前版本: %s
		联系方式:微信号：le_muzhi QQ：1163648924
		官网:https://www.yizhengtong.net
		本地接口地址:%v
	`+"\n", version, global.Config.GetString("gin.addr"))
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
	return nil
}
