package config

import (
	"strconv"
	"time"
)

// DBConfig 数据库相关配置
var DBConfig dbConfig

type dbConfig struct {
	Host               string
	Port               string
	Database           string
	User               string
	Password           string
	Charset            string
	MaxIdleConnections int
	MaxOpenConnections int
	ConnMaxLifetime    time.Duration
}

func init() {
	loadEnvFile()
	DBConfig.Host = DBHost
	DBConfig.Port = DBPort
	DBConfig.Database = DBDatabase
	DBConfig.User = DBUser
	DBConfig.Password = DBPassword
	DBConfig.Charset = DBCharset
	maxIdleConnections, err := strconv.Atoi(DBMaxOpenConnections)
	if err == nil {
		DBConfig.MaxIdleConnections = maxIdleConnections
	} else {
		DBConfig.MaxIdleConnections = 100
	}

	maxOpenConnections, err := strconv.Atoi(DBMaxOpenConnections)
	if err == nil {
		DBConfig.MaxOpenConnections = maxOpenConnections
	} else {
		DBConfig.MaxOpenConnections = 20
	}

	connMaxLifetime, err := strconv.Atoi(DBMaxOpenConnections)
	if err == nil {
		DBConfig.ConnMaxLifetime = time.Duration(connMaxLifetime) * time.Second
	} else {
		DBConfig.ConnMaxLifetime = 100 * time.Second
	}
}