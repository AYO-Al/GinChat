package router

import (
	"GinChat/docs"
	"GinChat/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Router() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	// 导入swag路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.GET("/swa")
	r.GET("/index", service.Getindex)
	r.GET("/user/data", service.GetData)
	r.GET("/user/create", service.CreateUser)
	r.GET("/user/delete", service.DeleteUser)
	return r
}
