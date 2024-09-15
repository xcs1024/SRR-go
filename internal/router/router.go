package router

import (
	"github.com/gin-gonic/gin"
	"reservation/internal/controller"
)

func Ping() *gin.Engine {
	ginSever := gin.Default()
	ginSever.GET("/ping", controller.PingHandler)
	return ginSever
}
