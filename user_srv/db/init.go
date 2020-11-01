package db

import (
	"fmt"
	"log"
	"user_srv/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var client *gorm.DB

var (
	username = config.DBConfig.User
	password = config.DBConfig.Password
	host     = config.DBConfig.Host
	port     = config.DBConfig.Port
	database = config.DBConfig.Database
	charset  = config.DBConfig.Charset
)

// 初始化链接
func init() {
	var err error
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)
	client, err = gorm.Open("mysql", dbDSN) // 打开连接失败
	if err != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + err.Error())
	}

	if gin.Mode() != "release" { // 非正式环境的数据库的调试log开关打开
		client.LogMode(true)
	}

	client.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConnections) // 最大连接数
	client.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConnections) // 闲置连接数
	client.DB().SetConnMaxLifetime(config.DBConfig.ConnMaxLifetime) // 最大连接周期
}

// Client for gorm.DB
func Client() *gorm.DB {
	return client
}

// Close for DB client
func Close() error {
	return client.Close()
}

// SearchByIDOrByName search by ID or search by Name
func SearchByIDOrByName(c *gin.Context, query *gorm.DB) *gorm.DB {
	searchByID := c.Query("search_by_id")
	searchByName := c.Query("search_by_name")

	if searchByID != "" {
		query = query.Where("id like ?", fmt.Sprintf("%%%v%%", searchByID))
	} else if searchByName != "" {
		query = query.Where("name like ?", fmt.Sprintf("%%%v%%", searchByName))
	}
	return query
}
