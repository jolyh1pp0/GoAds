package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type config struct {
	Database struct {
		Dialect  string
		Host     string
		Port     string
		User     string
		Name     string
		Password string
	}
	Server struct {
		Address string
	}
}

var C config

func ReadConfig() {
	Config := &C
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
