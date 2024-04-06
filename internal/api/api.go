package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var k = koanf.New(".")

type config struct {
	DB   string `koanf:"database.mysql"`
	Port string `koanf:"net.port"`
}

func StartServer() {
	// if err := k.Load(file.Provider(".env"), dotenv.Parser()); err != nil {
	// 	log.Fatalf("error loading config file: %v", err)
	// }

	if err := k.Load(env.Provider("dfogg_", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(s, "dfogg_")), "_", ".")
	}), nil); err != nil {
		log.Fatalf("error loading config from environment: %v", err)
	}

	log.Println("port", k.Get("net.port"))
	log.Println("name", k.Get("net.name"))
	log.Println("db", k.Get("database.mysql"))
	log.Println("convert", k.Int("net.port"))
	// k.Print()
	log.Println(k.All())

	var configuration config
	if err := k.UnmarshalWithConf("", &configuration, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true}); err != nil {
		log.Fatalf("error unmarshalling: %v", err)
	}
	log.Printf("%+v", configuration)
	log.Println(configuration)

	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		fmt.Println("error collected here")
	}
	adminRoutes := e.Group("/admin")
	adminRoutes.Use(middleware.KeyAuth(func(auth string, c echo.Context) (bool, error) {
		return auth == "test", nil
	}))
	adminRoutes.GET("/stats", func(c echo.Context) error {
		return c.String(http.StatusOK, "data here")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
