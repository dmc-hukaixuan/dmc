package database

import (
	"dmc/config/database"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db2 *gorm.DB

func GormMysqlReport() *gorm.DB {
	username := database.GetConfig().MySQLReport.User     //账号
	password := database.GetConfig().MySQLReport.Password //密码
	host := database.GetConfig().MySQLReport.Host         //数据库地址，可以是Ip或者域名
	port := database.GetConfig().MySQLReport.Port         //数据库端口
	Dbname := database.GetConfig().MySQLReport.Dbname     //数据库名
	timeout := "10s"                                      //连接超时，10秒
	fmt.Println("db port :", port, "host", host, "Dbname ", Dbname)
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	fmt.Println("db dsn :", dsn)
	var err error

	//连接 MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_db2, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := _db2.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	return _db2
}

// func GetDB() *gorm.DB {
// 	return _db2
// }
