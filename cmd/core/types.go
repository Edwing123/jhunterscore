package main

import (
	"log/slog"

	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Core struct {
	Logger   *slog.Logger
	Store    *session.Store
	Database models.Database
}

type Config struct {
	ServerPort uint16 `json:"server_port"`
	DataDir    string `json:"data_dir"`
}
