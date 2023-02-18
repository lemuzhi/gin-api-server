package global

import (
	"gin-project-template/pkg/logger"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Config *viper.Viper
	Log    *logger.Logger
	DB     *gorm.DB
	RDB    *redis.Client
)
