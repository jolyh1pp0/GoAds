package config

import (
	"GoAds/domain"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
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
	JWT struct {
		Key                  string
		RefreshKey           string
		AccessTokenLifespan  time.Duration
		RefreshTokenLifespan time.Duration
	}
	S3 struct {
		BucketName string
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

	if C.JWT.Key == "" || C.JWT.RefreshKey == ""{
		log.Println(domain.ErrEmptyJWTKey)
		os.Exit(1)
	}
}
