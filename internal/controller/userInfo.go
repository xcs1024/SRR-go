package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func GetUserInfo(context *gin.Context) {
	var userInfo UserInfo
	err := context.ShouldBindJSON(&userInfo)
	if err != nil {
		log.Println(err)
		context.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	context.JSON(200, gin.H{
		"message": "success",
		"data":    userInfo,
	})
}
