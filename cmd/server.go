package cmd

import (
	"context"
	"fmt"
	"gin-api-server/conf"
	"gin-api-server/initialize"
	"gin-api-server/internal"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
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
	cmd.PersistentFlags().StringVarP(&opts.config, "conf", "c", "./conf/conf.toml", "Config file")
	return cmd
}

func RunServerStart(ctx context.Context, opts *ServerStartOptions, version string) error {
	//初始化配置
	initialize.InitConfig("./conf/config.toml")
	//初始化日志
	initialize.InitLogger()
	zap.L().Info("heello")
	zap.L().Debug("heello")
	zap.L().Warn("heello")
	//初始化mysql
	initialize.InitMysql()

	//初始化redis
	initialize.InitRedis()

	//初始化定时任务
	initialize.InitCron()

	//设置gin的启动模式
	gin.SetMode(conf.Server.Mode)

	//初始化路由
	handler := internal.InitRouter()

	//自定义HTTP配置
	server := &http.Server{
		Addr:           conf.Server.Addr,
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
	`+"\n", version, conf.Server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
	return nil
}
