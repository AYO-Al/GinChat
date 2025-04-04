package models

import (
	"GinChat/app"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Avatar        string //头像
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUser() []*UserBasic {
	var users []*UserBasic
	app.Db.Find(&users)
	return users
}

func CreateUsers(user UserBasic) *gorm.DB {
	return app.Db.Create(&user)
}

func DeleteUsers(name string) *gorm.DB {
	var user UserBasic
	return app.Db.Where("name = ?", name).Delete(&user)
}
