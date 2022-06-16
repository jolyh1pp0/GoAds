package main

import (
	"GoAds/config"
	"GoAds/infrastructure/datastore"
	"GoAds/infrastructure/routes"
	"GoAds/registry"
	"fmt"
	"github.com/labstack/echo"
	"log"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = routes.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
