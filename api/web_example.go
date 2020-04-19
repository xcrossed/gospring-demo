package api

import (
	"fmt"
	"net/http"

	"github.com/xcrossed/gospring-demo/service"

	SpringWeb "github.com/go-spring/go-spring-web/spring-web"
)

type MyController struct {
	MyWebService *service.MyWebService `autowire:""`
}

func (c *MyController) Home(ctx SpringWeb.WebContext) {
	ctx.String(http.StatusOK, fmt.Sprintf("hello world."))
}

func (c *MyController) Echo(ctx SpringWeb.WebContext) {
	bizRet := c.MyWebService.Biz()
	hi := ctx.QueryParam("hi")
	ctx.String(http.StatusOK, fmt.Sprintf("OK,Biz return :%s,echo return:%s", bizRet, hi))
}
