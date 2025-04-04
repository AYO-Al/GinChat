package db

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
	"time"
)

var Db *gorm.DB

type dbConf struct {
	Addr     string `mapstructure:"addr"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

// json反序列化
func getConf() string {
	var dbconf = new(dbConf)

	f, err := os.Open("./conf/mysql_conf.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	conf, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(conf, dbconf)
	if err != nil {
		fmt.Println(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconf.User, dbconf.Password, dbconf.Addr, dbconf.Port, dbconf.Dbname)
	return dsn
}

// 使用viper获取配置
func VgetConf() string {
	var dbconf dbConf
	v := viper.New()
	v.SetConfigFile("conf/mysql_conf.json")
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(err)
		} else {
			fmt.Println(err)
		}
	}

	err = v.Unmarshal(&dbconf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dbconf)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconf.User, dbconf.Password, dbconf.Addr, dbconf.Port, dbconf.Dbname)
	return dsn
}

func init() {
	dsn := VgetConf()
	fmt.Println(dsn)

	dbc, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Db = dbc
	db, err := Db.DB()
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
}
