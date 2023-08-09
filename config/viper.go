package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

type envConfigs struct {
	DB_USER string
	DB_PASS string
	DB_NAME string
	DB_PORT int64
}

func loadEnvVariables() (config *envConfigs) {

	workingdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile(workingdir + "/.env")

	// read global vars from env
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// unmarshal into struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
