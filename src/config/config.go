package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Label struct {
	Text     string
	X        float64
	Y        float64
	FontSize float64
	Font     string
	Color    string
}

type Connection struct {
	Host     string
	Port     int
	ServerId int
	User     string
	Password string
}
type Config struct {
	RefreshInterval int
	Connection      Connection
	Labels          []Label
	TemplatePath    string
}

func LoadConfig() Config {
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.SetDefault("RefreshInterval", 60)
	viper.SetDefault("TemplatePath", "template.png")

	viper.SetDefault("Connection.Port", 10011)
	viper.SetDefault("Connection.ServerId", 1)
	viper.SetDefault("Connection.User", "serveradmin")

	var C Config

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.Unmarshal(&C)
	return C
}
