package config

import (
	"GoAds/domain"
	"context"
	"fmt"
	cfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mailgun/mailgun-go/v4"
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
	Mailgun struct {
		DomainName string
		ApiKey     string
	}
}

var C config

var BucketClient *s3.Client

var MG *mailgun.MailgunImpl

func createBucket() (*s3.Client, error) {
	bucketConfig, err := cfg.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
		return nil, err
	}

	client := s3.NewFromConfig(bucketConfig)

	return client, nil
}

func ReadConfig(configPath, configName, configType string) {
	Config := &C

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

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

	var err error
	BucketClient, err = createBucket()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	MG = mailgun.NewMailgun(C.Mailgun.DomainName, C.Mailgun.ApiKey)
}
