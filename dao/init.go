package dao

import (
	"github.com/xcrossed/gospring-demo/dao/mysql"
	"github.com/xcrossed/gospring-demo/dao/redis"
	"github.com/xcrossed/gospring-demo/dao/user"
)

func Init() {
	mysql.Init()
	user.Init()

	redis.Init()
}
