package initialize

import (
	"context"
	"gin-project-template/global"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
)

var (
	ctx = context.Background()
)

func InitRedis() {
	//fmt.Println(viper.GetString("redis.url"))
	//opt, err := redis.ParseURL(viper.GetString("redis.url"))
	//if err != nil {
	//	log.Println("rdb.Ping(ctx).Result() = ", err)
	//}
	//rdb := redis.NewClient(opt)
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Username: viper.GetString("redis.username"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("rdb.Ping(ctx).Result() = ", err)
	}
	global.RDB = rdb
}
