package tests

import (
	"GoAds/config"
	"GoAds/infrastructure/datastore"
	"GoAds/infrastructure/routes"
	"GoAds/registry"
	"GoAds/tests/test_db"
	"github.com/labstack/echo"
	"os"
	"testing"
)

var EchoServer *echo.Echo

func TestMain(m *testing.M) {
	config.ReadConfig("./test_config", "config_test", "yml")

	EchoServer = echo.New()
	db := datastore.NewDB()
	defer db.Close()

	test_db.ResetDB()

	r := registry.NewRegistry(db)
	EchoServer = routes.NewRouter(EchoServer, r.NewAppController())

	code := m.Run()
	os.Exit(code)
}
