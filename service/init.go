package service

import SpringBoot "github.com/go-spring/go-spring/spring-boot"

func Init() {
	SpringBoot.RegisterBean(new(MyWebService))
	SpringBoot.RegisterBean(new(MyRedisService))
}
