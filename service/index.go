package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Getindex(c *gin.Context) {
	c.String(http.StatusOK, "welcome !!")
}
