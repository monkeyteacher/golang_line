package configs

import (
	"log"

	"github.com/spf13/viper"
)

// Initilize this variable to access the env values
var EnvConfigs *envConfigs

// We will call this in main.go to load the env variables
func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

// struct to map env values
type envConfigs struct {
	ServerProt        string `mapstructure:"SERVER_PROT"`
	LineChannelSecret string `mapstructure:"CHANNEL_SECRET"`
	LineChannelToken  string `mapstructure:"CHANNEL_TOKEN"`
	DBIP              string `mapstructure:"DB_IP"`
	DBName            string `mapstructure:"DB_NAME"`
}

// Call to load the variables from env
func loadEnvVariables() (config *envConfigs) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath(".")

	// Tell viper the name of your file
	viper.SetConfigName("app")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
