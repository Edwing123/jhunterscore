package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func (core *Core) GetSession(c *fiber.Ctx) *session.Session {
	return c.Locals(SESSION_KEY).(*session.Session)
}

func GetConfig() Config {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	if configPath == nil || *configPath == "" {
		panic("Missing config file path")
	}

	configFile, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	var config Config

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
func InitDataDir(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(path+"/"+DATA_DIR_LOGS_DIR, 0755)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(path+"/"+DATA_DIR_SESSION_DIR, 0755)
	if err != nil {
		panic(err)
	}
}
