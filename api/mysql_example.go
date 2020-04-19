package api

import (
	"fmt"
	"net/http"
	"strconv"

	SpringLogger "github.com/go-spring/go-spring-parent/spring-logger"
	SpringWeb "github.com/go-spring/go-spring-web/spring-web"
	"github.com/xcrossed/gospring-demo/dao/objects"
	"github.com/xcrossed/gospring-demo/service"
)

type MySQLController struct {
	MySQLService *service.MySQLService `autowire:""`
}

func (c *MySQLController) RegisterUser(ctx SpringWeb.WebContext) {
	username := ctx.QueryParam("username")
	age := ctx.QueryParam("age")
	// age to int
	ageInt, err := strconv.ParseUint(age, 10, 64)
	if err != nil {
		ctx.String(http.StatusOK, fmt.Sprintf("parameter age:%v is err:%v", ageInt, err))
	}

	user := &objects.User{Username: username, Age: ageInt}
	err = c.MySQLService.SaveUser(user)
	if err != nil {
		SpringLogger.Errorf("err:%v", err)
		ctx.String(http.StatusOK, fmt.Sprintf("ce.SaveUser(%v) is err:%v", user, err))
	}

	ctx.String(http.StatusOK, "ok")
}
