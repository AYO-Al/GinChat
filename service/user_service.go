package service

import (
	"GinChat/models"
	"GinChat/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func GetData(c *gin.Context) {
	c.String(http.StatusOK, "user_data:%s", models.GetUser())
}

// CreateUser
// @Tags 用户模块
// @Summary 新增用户
// @param name formData string false "name"
// @param password formData string false "password"
// @param email formData string false "email"
// @param phone formData string false "phone"
// @Success 200 {string} json{"code","msg"}
// @Router /user/create [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	if models.FindUserByPhone(user.Phone) || models.FindUserByEmail(user.Email) || models.FindUserByName(user.Name) {
		c.JSON(-1, gin.H{
			"msg": "用户名/电话/邮箱已存在",
		})
		return
	}

	db := models.CreateUsers(user)

	if db.Error != nil {
		c.JSON(-1, gin.H{
			"msg": db.Error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": fmt.Sprintf("新增用户成功:%s", user),
		})
	}
}

// DeleteUser
// @Tags 用户模块
// @Summary 删除用户
// @param name query string false "name"
// @Success 200 {string} json{"code","msg"}
// @Router /user/delete [get]
func DeleteUser(c *gin.Context) {

	name := c.Query("name")
	db := models.DeleteUsers(name)

	if db.Error != nil {
		c.JSON(-1, gin.H{
			"msg": db.Error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": fmt.Sprintf("删除用户%s成功", name),
		})
	}
}

// UpdateUser
// @Tags 用户模块
// @Summary 修改用户
// @param name formData string false "name"
// @param update_name formData string false "update_name"
// @param update_password formData string false "update_password"
// @param update_email formData string false "update_email"
// @param update_phone formData string false "update_phone"
// @Success 200 {string} json{"code","msg"}
// @Router /user/update [post]
func UpdateUser(c *gin.Context) {
	var user models.UserBasic
	validate := validator.New()

	name := c.PostForm("name")
	user.Name = c.PostForm("update_name")
	user.PassWord = c.PostForm("update_password")
	user.Email = c.PostForm("update_email")
	user.Phone = c.PostForm("update_phone")

	// 注册自定义验证器
	validate.RegisterValidation("phone", utils.ValidaPhone)

	if err := validate.Struct(user); err != nil {
		c.JSON(-1, gin.H{
			"msg": err.Error(),
		})
		return
	}
	db := models.UpdateUsers(name, user)

	if db.Error != nil {
		c.JSON(-1, gin.H{
			"msg": db.Error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": fmt.Sprintf("修改用户%s后信息:\nname:%s\npassword:%s\n", name, user.Name, user.PassWord),
		})
	}
}
