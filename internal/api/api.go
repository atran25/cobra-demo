package api

import (
	"fmt"
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

func StartServer() {
	if err := k.Load(file.Provider("./.env"), dotenv.Parser()); err != nil {
		log.Fatalf("error loading config file: %v", err)
	}

	if err := k.Load(env.Provider("dfogg_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "dfogg_")), "_", ".", -1)
	}), nil); err != nil {
		log.Fatalf("error loading config from environment: %v", err)
	}

	fmt.Println("istmall", k.Get("istmall"))
	fmt.Println("istmall2", k.Get("istmall2"))
}
