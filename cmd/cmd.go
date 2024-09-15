package cmd

import (
	"log"
	"reservation/internal/router"
)

func RunServer() {
	ginSever := router.Ping()
	err := ginSever.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
