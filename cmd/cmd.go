package cmd

import (
	"log"
	"reservation/internal/database"
	"reservation/internal/router"
)

func RunServer() {
	database.MysqlConnect()
	ginSever := router.Ping()
	ginSever = router.GetUserInfo()
	err := ginSever.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
