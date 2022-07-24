package main

import (
	"GoAds/config"
	"GoAds/infrastructure/datastore"
	"GoAds/infrastructure/routes"
	"GoAds/registry"
	"fmt"
	"github.com/labstack/echo"
	logrus "github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	logrus.SetLevel(logLevel)
}

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()

	e.Use(loggingMiddleware)

	e = routes.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}

func loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		res := next(c)

		logrus.WithFields(logrus.Fields{
			"method":     c.Request().Method,
			"path":       c.Path(),
			"status":     c.Response().Status,
			"latency_ns": time.Since(start).Nanoseconds(),
			"user_ip":    c.Request().RemoteAddr,
		}).Info("request details")

		return res
	}
}
