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
	Phone         string `validate:"phone"`
	Email         string `validate:"email"`
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

func UpdateUsers(name string, user UserBasic) *gorm.DB {
	return app.Db.Model(&user).Where("name = ?", name).Updates(&user)
}
