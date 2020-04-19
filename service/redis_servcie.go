package service

import (
	"time"

	SpringLogger "github.com/go-spring/go-spring-parent/spring-logger"
	"gopkg.in/redis.v5"
)

type MyRedisService struct {
	RedisClient redis.Cmdable `autowire:"redis.default"`
}

func (s *MyRedisService) Get(key string) (string, error) {
	val := s.RedisClient.Get(key)
	if err := val.Err(); err != nil && err != redis.Nil {
		SpringLogger.Errorf("get key:%s, err:%v", key, err)
		return "", err
	}
	return val.Val(), nil
}

func (s *MyRedisService) Set(key, val string) error {
	err := s.RedisClient.Set(key, val, time.Duration(1000)*time.Second).Err()
	if err != nil {
		SpringLogger.Errorf("set key:%s, err:%v", key, err)
		return err
	}
	return nil
}
