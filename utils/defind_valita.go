package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

// 自定义验证器
var ValidaPhone validator.Func = func(fl validator.FieldLevel) bool {
	phone := fl.Field().Interface().(string)
	if len(phone) != 11 {
		fmt.Println("Phone is len error")
		return false
	} else if ok, _ := regexp.MatchString("^1[3-9]{1}\\d{9}$", phone); !ok {
		fmt.Println("Phone is error")
		return false
	} else {
		return true
	}
}
