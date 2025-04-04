package service

import (
	"GinChat/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetData(c *gin.Context) {
	c.String(http.StatusOK, "user_data:%s", models.GetUser())
}

// CreateUser
// @Tags 用户模块
// @Summary 新增用户
// @Success 200 {string} json{"code","msg"}
// @Router /user/create [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	user.PassWord = c.Query("password")
	db := models.CreateUsers(user)

	if db.Error != nil {
		c.JSON(-1, gin.H{
			"msg": db.Error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "新增用户成功",
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
