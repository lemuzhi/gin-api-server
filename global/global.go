package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Config *viper.Viper
	DB     *gorm.DB
	RDB    *redis.Client
)
