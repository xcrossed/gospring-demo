package config

type WebServerConf struct {
	Addr string `value:"${web.server.address}"`
	Port int    `value:"${web.server.port}"`
}

type MySQLDataSourceConf struct {
	Driver             string `value:"${db.datasource.driver}"`
	Dsn                string `value:"${db.datasource.dsn}"`
	MaxOpenConnections int    `value:"${db.datasource.maxOpenConnections}"`
	MaxIdleConnections int    `value:"${db.datasource.maxIdleConnections}"`
	MaxLifeTimeSeconds int    `value:"${db.datasource.maxLifeTimeSeconds}"`
}

type RedisDataSourceConf struct {
	Addr string `value:"${redis.datasource.address}"`
	Pwd  string `value:"${redis.datasource.password}"`
	Db   int    `value:"${redis.datasource.db}"`
}
