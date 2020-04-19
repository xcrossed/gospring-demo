package service

import (
	SpringLogger "github.com/go-spring/go-spring-parent/spring-logger"
	"github.com/xcrossed/gospring-demo/dao/objects"
	"github.com/xcrossed/gospring-demo/dao/user"
)

type MySQLService struct {
	UserDao *user.UserDao `autowire:"user.dao"`
}

func (myService *MySQLService) SaveUser(user *objects.User) error {
	_, err := myService.UserDao.Insert(user)
	if err != nil {
		SpringLogger.Errorf("SaveUser(%v),err:%v", user, err)
		return err
	}
	return err
}
