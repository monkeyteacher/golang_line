package main

import (
	"golang_line/configs"
)

func main() {
	configs.InitEnvConfigs()
	// fmt.Println(configs.EnvConfigs.LineChannelSecret)
	// fmt.Println(configs.EnvConfigs.LineChannelToken)
}
