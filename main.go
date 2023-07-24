package main

import (
	"golang_line/configs"
	"golang_line/database"
	"golang_line/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.InitEnvConfigs()
	database.DBInit()
	mainServer := gin.Default()
	routes.ApiRoutes(mainServer)
	mainServer.Run(":8088")
}
