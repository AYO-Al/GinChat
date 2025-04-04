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

func FindUserByName(name string) bool {
	if app.Db.Where("name = ?", name).Find(&UserBasic{}).RowsAffected != 0 {
		return true
	} else {
		return false
	}
}

func FindUserByPhone(phone string) bool {
	if app.Db.Where("phone = ?", phone).Find(&UserBasic{}).RowsAffected != 0 {
		return true
	} else {
		return false
	}
}

func FindUserByEmail(email string) bool {
	if app.Db.Where("email = ?", email).Find(&UserBasic{}).RowsAffected != 0 {
		return true
	} else {
		return false
	}
}

func FindName(user *UserBasic) bool {
	app.Db.Where("name = ?", user.Name).First(&user)
	if user.Salt != "" {
		return true
	} else {
		return false
	}
}
