package app

import (
	"GinChat/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func init() {

	// 日志设置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // 日志等级
			Colorful:      true,        // 彩色
		},
	)

	dbc, err := gorm.Open(mysql.Open(utils.VgetConf()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 如果直接在上面赋值，会覆盖全局变量
	Db = dbc
	db, err := Db.DB()
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
}
