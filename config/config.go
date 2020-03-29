package config

type WebServerConf struct {
	Addr string `value:"${web.server.address}"`
	Port int    `value:"${web.server.port}"`
}

type MySQLDataSourceConf struct {
	Dsn                string `value:"${mysql.datasource.dsn}"`
	MaxOpenConnections int    `value:"${mysql.datasource.maxOpenConnections}"`
	MaxIdleConnections int    `value:"${mysql.datasource.maxIdleConnections}"`
	MaxLifeTime        int    `value:"${mysql.datasource.maxLifeTime}"`
}

type RedisDataSourceConf struct {
	Addr string `value:"${redis.datasource.address}"`
	Pwd  string `value:"${redis.datasource.password}"`
	Db   int    `value:"${redis.datasource.db}"`
}
