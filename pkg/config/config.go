package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	DBUrl        string `mapstructure:"DB_URL"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (Config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("env")
	viper.SetConfigFile("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println("env config error")
		return
	}

	err = viper.Unmarshal(&Config)
	return

}
