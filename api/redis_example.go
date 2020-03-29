package api

import (
	"fmt"
	"net/http"

	"github.com/xcrossed/gospring-demo/service"

	SpringWeb "github.com/go-spring/go-spring-web/spring-web"
)

type MyRedisController struct {
	MyRedisService *service.MyRedisService `autowire:""`
}

func (c *MyRedisController) Get(ctx SpringWeb.WebContext) {
	key := ctx.QueryParam("key")
	val, _ := c.MyRedisService.Get(key)
	ctx.String(http.StatusOK, fmt.Sprintf("OK,get return :%v", val))
}

func (c *MyRedisController) Set(ctx SpringWeb.WebContext) {
	key := ctx.QueryParam("key")
	val := ctx.QueryParam("val")
	err := c.MyRedisService.Set(key, val)
	ctx.String(http.StatusOK, fmt.Sprintf("OK,set return:%v", err))
}
