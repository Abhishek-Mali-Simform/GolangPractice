package utils

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

type envConfigs struct {
	DBHost     string `mapstructure:"HOST"`
	DBPort     string `mapstructure:"DBPORT"`
	DBUser     string `mapstructure:"USER"`
	DBName     string `mapstructure:"NAME"`
	DBPassword string `mapstructure:"PASSWORD"`
	DBAlias    string `mapstructure:"ALIAS"`
}

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *envConfigs) {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		logs.Error("Error Reading Config File. Reason: ", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		logs.Error("Error Unmarshalling Config File. Reason: ", err)
	}
	return
}
