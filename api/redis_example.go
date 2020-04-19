package api

import (
	"fmt"
	"net/http"

	"github.com/xcrossed/gospring-demo/service"

	SpringLogger "github.com/go-spring/go-spring-parent/spring-logger"
	SpringWeb "github.com/go-spring/go-spring-web/spring-web"
)

type MyRedisController struct {
	MyRedisService *service.MyRedisService `autowire:""`
}

func (c *MyRedisController) Get(ctx SpringWeb.WebContext) {
	key := ctx.QueryParam("key")
	val, err := c.MyRedisService.Get(key)
	if err != nil {
		SpringLogger.Errorf("MyRedisService.Get(%v) error: %v", key, err)
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("error,get return :%v", err))
		return
	}
	ctx.String(http.StatusOK, fmt.Sprintf("OK,get return :%v", val))
}

func (c *MyRedisController) Set(ctx SpringWeb.WebContext) {
	key := ctx.QueryParam("key")
	val := ctx.QueryParam("val")
	err := c.MyRedisService.Set(key, val)
	ctx.String(http.StatusOK, fmt.Sprintf("OK,set return:%v", err))
}
