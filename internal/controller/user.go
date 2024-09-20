package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"reservation/internal/database"
	"reservation/internal/model"
)

// GetUser 根据name查询一条记录
func GetUser(context *gin.Context) {
	var user model.User
	err := database.DB.Take(&user, "name = ?", "机器人27号").Error
	if err != nil {
		log.Fatalln("GetUser+", err)
	}
	context.JSON(200, gin.H{
		"message": "success",
		"data":    user,
	})
}

// CreateUser 批量插入
func CreateUser(context *gin.Context) {
	var userList []model.User
	for i := 0; i < 100; i++ {
		userList = append(userList, model.User{
			Name:    fmt.Sprintf("机器人%d号", i+1),
			Age:     20 + i,
			Sex:     1,
			Address: fmt.Sprintf("机器人%d号地址", i+1),
		})
	}
	err := database.DB.Create(&userList)
	if err != nil {
		log.Fatalln("CreateUser+", err)
	}
	{
		context.JSON(200, gin.H{
			"message": "success",
			"data":    "ok",
		})
	}
}

// GetUserListById 根据主键批量查询
func GetUserListById(context *gin.Context) {
	var userList []model.User
	err := database.DB.Find(&userList, []int{1, 3, 4, 7}).Error
	if err != nil {
		log.Fatalln("GetUserListById+", err)
	} else {
		context.JSON(200, gin.H{
			"message": "success",
			"data":    userList,
		})
	}
}

// GetUserListByName 根据name批量查询
func GetUserListByName(context *gin.Context) {
	var userList []model.User
	err := database.DB.Find(&userList, "name in?", []string{"机器人1号", "机器人2号", "机器人3号", "机器人4号"}).Error
	if err != nil {
		log.Fatalln("GetUserListByName+", err)
	} else {
		context.JSON(200, gin.H{
			"message": "success",
			"data":    userList,
		})
	}
}
