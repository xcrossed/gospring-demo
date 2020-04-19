package main

import (
	"github.com/xcrossed/gospring-demo/dao"
	"github.com/xcrossed/gospring-demo/service"

	"github.com/xcrossed/gospring-demo/api"

	SpringLogger "github.com/go-spring/go-spring-parent/spring-logger"
	SpringBoot "github.com/go-spring/go-spring/spring-boot"
	_ "github.com/go-spring/go-spring/starter-gin"
)

func main() {
	dao.Init()
	service.Init()
	api.Init()
	level := SpringLogger.InfoLevel
	console := &SpringLogger.Console{}
	console.SetLevel(level)
	SpringLogger.SetLogger(console)
	SpringBoot.RunApplication("config/")
}
