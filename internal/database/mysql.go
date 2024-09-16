package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

const (
	mysqlHost     = "127.0.0.1"
	mysqlPort     = "3306"
	mysqlUser     = "root"
	mysqlPassword = "123456"
	mysqlDB       = "reservation_dev"
	timeout       = "10s"
)

func MysqlConnect() {
	//gorm的默认日志是只打印错误和慢Sql
	var mysqlLogger logger.Interface
	//要显示的日志等级
	mysqlLogger = logger.Default.LogMode(logger.Info)

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDB, timeout)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		Logger: mysqlLogger, //打印日志
		//SkipDefaultTransaction: true, // 禁用默认事务(为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除))
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix:   "t_",  //表名前缀
		//	SingularTable: true,  // 禁用表名复数
		//	NoLowerCase:   false, // 禁用小写
		//},
	})
	if err != nil {
		log.Fatalln(err)
	}
	//连接成功
	DB = db
	log.Println(db)
}
