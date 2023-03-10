package initialize

import (
	"context"
	"gin-api-server/conf"
	"gin-api-server/global"
	"github.com/go-redis/redis/v8"
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
		Addr:     conf.Redis.Host,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("rdb.Ping(ctx).Result() = ", err)
	}
	global.RDB = rdb
}
