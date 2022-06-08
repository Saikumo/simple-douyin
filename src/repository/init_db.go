package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	. "saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/entity"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(Config.GetString("database.dsn")), &gorm.Config{
		PrepareStmt:            true,                                //缓存预编译命令
		SkipDefaultTransaction: true,                                //禁用默认事务操作
		Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		panic(err)
	}

	//自动建表
	if err := DB.AutoMigrate(&entity.User{}); err != nil {
		panic(err)
	}
}
