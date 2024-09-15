package main

import (
	"fmt"
	"lecter/hello/common"
	"lecter/hello/db"
	"lecter/hello/router"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// configs.jsonから読み込んだConfig
	appConfig := common.LoadConfig()
	if (appConfig != nil) {
		common.ApplicationConfig = *appConfig
	}

	// DB接続
	err := db.Connect()
	if (err != nil) {
		fmt.Println("DB Connection ERROR")
		return
	}
	defer db.Database().Close()

	server := gin.Default()
	router.Routing(server)

	// サーバー起動
	server.Run(":" + strconv.Itoa(common.ApplicationConfig.Port))
}