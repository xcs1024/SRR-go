package cmd

import "log"

func RunServer() {
	ginSever := router.New()
	err := ginSever.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
