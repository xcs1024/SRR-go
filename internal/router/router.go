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

func UserHandler() *gin.Engine {
	ginSever := gin.Default()
	r := ginSever.Group("/user")
	r.GET("/getUser", controller.GetUser)
	r.GET("/createUser", controller.CreateUser)
	r.GET("/getUserById", controller.GetUserListById)
	r.GET("/getUserByName", controller.GetUserListByName)
	return ginSever
}
