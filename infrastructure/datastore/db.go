package datastore

import (
	"GoAds/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func NewDB() *gorm.DB {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", config.C.Database.Host, config.C.Database.User, config.C.Database.Name, config.C.Database.Password, config.C.Database.Port)
	db, err := gorm.Open(config.C.Database.Dialect, dbURI)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connected to PostgreSQL!")
	}

	return db
}
