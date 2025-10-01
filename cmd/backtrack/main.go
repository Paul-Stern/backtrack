package main

import (
	"os"
	"path/filepath"
	"time"

	cmd "github.com/Paul-Stern/backtrack/commands"
	"github.com/Paul-Stern/backtrack/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	InitLogger()
	model.InitDB()
	cmd.Execute()
}
func InitLogger() {

	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err = os.Mkdir(logDir, 0755); err != nil {
			panic(err)
		}
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Create file writer with rotation
	fileWriter := &lumberjack.Logger{

		Filename:   filepath.Join(logDir, "app.log"),
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, fileWriter)

	logger := zerolog.New(multi).With().Timestamp().Logger()
	log.Logger = logger
}
