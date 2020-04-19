package mysql

import (
	"fmt"
	"time"

	SpringLogger "github.com/go-spring/go-spring-parent/spring-logger"
	SpringBoot "github.com/go-spring/go-spring/spring-boot"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xcrossed/gospring-demo/config"
)

const (
	DefaultDriver           string        = "mysql"
	MySQLMaxOpenConnections int           = 15
	MySQLMaxIdleConnections int           = 20
	MySQLMaxLifeTime        time.Duration = time.Minute * 10
)

func Init() {
	SpringBoot.RegisterNameBeanFn("mysql.default", NewMySqlClient).
		ConditionOnMissingBean((*sqlx.DB)(nil))
}

func NewMySqlClient(conf config.MySQLDataSourceConf) (*sqlx.DB, error) {
	if conf.MaxIdleConnections <= 0 {
		conf.MaxIdleConnections = MySQLMaxIdleConnections
	}
	if conf.MaxOpenConnections <= 0 {
		conf.MaxOpenConnections = MySQLMaxOpenConnections
	}
	maxLifeTimeSeconds := MySQLMaxLifeTime
	if conf.MaxLifeTimeSeconds > 0 {
		maxLifeTimeSeconds = time.Second * time.Duration(conf.MaxLifeTimeSeconds)
	}
	// 此处drivername为mysql，可以根据实际需求换成pg等其它driver
	db, err := sqlx.Connect(conf.Driver, conf.Dsn)
	if err != nil {
		SpringLogger.Errorf("open mysql err:%v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	name := "unknown>"
	if cfg, err := mysql.ParseDSN(conf.Dsn); err == nil {
		name = fmt.Sprintf("%s/%s", cfg.Addr, cfg.DBName)
	}

	SpringLogger.Infof("connect to mysql [%v] success", name)

	db.SetMaxOpenConns(conf.MaxOpenConnections)
	db.SetMaxIdleConns(conf.MaxIdleConnections)
	db.SetConnMaxLifetime(maxLifeTimeSeconds)

	return db, nil
}
