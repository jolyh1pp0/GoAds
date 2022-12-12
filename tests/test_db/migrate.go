package test_db

import "C"
import (
	"GoAds/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"log"
)

func ResetDB() {
	DSN := "postgres://" + config.C.Database.Dialect + ":" + config.C.Database.Password + "@" +
		config.C.Database.Host + ":" + config.C.Database.Port + "/" + config.C.Database.Name +
		"?sslmode=disable"

	M, err := migrate.New("file:///Code/GoAds/tests/test_db/test_migrations", DSN)
	if err != nil {
		log.Print(err)
	}

	err = M.Down()
	if err != nil {
		log.Print(err)
	}

	err = M.Up()
	if err != nil {
		log.Print(err)
	}
}
