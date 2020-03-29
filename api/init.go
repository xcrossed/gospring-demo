package api

import (
	SpringBoot "github.com/go-spring/go-spring/spring-boot"
)

func Init() {
	// router registe
	SpringBoot.RegisterBean(new(MyController)).Init(func(c *MyController) {
		SpringBoot.GetMapping("/", c.Home)
		SpringBoot.GetMapping("/echo", c.Echo)
	})

	SpringBoot.RegisterBean(new(MyRedisController)).Init(func(c *MyRedisController) {
		SpringBoot.GetMapping("/get", c.Get)
		SpringBoot.GetMapping("/set", c.Set)
	})
}
