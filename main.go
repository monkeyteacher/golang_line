package main

import (
	"golang_line/configs"
	"golang_line/database"
)

func main() {
	configs.InitEnvConfigs()
	database.DBInit()
}
