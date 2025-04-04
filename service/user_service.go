package service

import (
	"GinChat/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetData(c *gin.Context) {
	c.String(http.StatusOK, "user_data:%s", models.GetUser())
}
