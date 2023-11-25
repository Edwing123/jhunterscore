package main

import (
	"fmt"
	"os"
)

func main() {
	config := GetConfig()

	DATA_DIR_LOGS_PATH := config.DataDir + "/" + DATA_DIR_LOGS_DIR + "/" + LOGS_FILE
	DATA_DIR_SESSION_PATH := config.DataDir + "/" + DATA_DIR_SESSION_DIR + "/" + SESSION_FILE

	InitDataDir(config.DataDir)

	logFile, err := os.OpenFile(DATA_DIR_LOGS_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logger := NewLogger(logFile)
	sessionStore := NewSessionStore(DATA_DIR_SESSION_PATH)

	core := Core{
		Logger: logger,
		Store:  sessionStore,
	}

	app := core.Setup()
	app.Listen(fmt.Sprintf(":%d", config.ServerPort))
}
