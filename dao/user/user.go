package user

import (
	SpringBoot "github.com/go-spring/go-spring/spring-boot"
	"github.com/jmoiron/sqlx"
	"github.com/xcrossed/gospring-demo/dao/objects"
)

type IUserDao interface {
	Insert(user *objects.User) (int64, error)
}

var _ IUserDao = &UserDao{}

type UserDao struct {
	DB *sqlx.DB `autowire:"mysql.default"`
}

func Init() {
	SpringBoot.RegisterNameBean("user.dao", &UserDao{})
}
func (dao *UserDao) Insert(user *objects.User) (int64, error) {
	// dao.DB.Exec("insert into users(username,password,email,age,sex) values(?,?,?,?,?)")
	sql := "insert into users(username,age) values(?,?)"
	args := []interface{}{user.Username, user.Age}
	result, err := dao.DB.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}
