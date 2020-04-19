package redis

import (
	SpringLogger "github.com/go-spring/go-spring-parent/spring-logger"
	SpringBoot "github.com/go-spring/go-spring/spring-boot"
	"github.com/xcrossed/gospring-demo/config"
	"gopkg.in/redis.v5"
)

func Init() {
	SpringBoot.RegisterNameBeanFn("redis.default", NewGoRedisClient).
		ConditionOnMissingBean((*redis.Cmdable)(nil))
}

// NewGoRedisClient 创建 redis 客户端
func NewGoRedisClient(config config.RedisDataSourceConf) (redis.Cmdable, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Pwd,
		DB:       config.Db,
	})
	if err := client.Ping().Err(); err != nil {
		SpringLogger.Errorf("redis connect failed: %s", err)
		return nil, err
	}
	SpringLogger.Infof("connect to redis [%v] success", config.Addr)
	return client, nil
}
