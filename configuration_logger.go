package configuration

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
)

var logger *logrus.Logger

func initLogger() {
	logFolder := viper.GetString("log.folder")
	err := os.MkdirAll(logFolder, os.ModePerm)
	if err != nil {
		fmt.Errorf("error creating log folder: %v", err)
		os.Exit(1)
	}
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetLevel(getLogLevel())

	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logFolder, "configuration.log"),
		MaxSize:    GetIntOrDefault("log.file.maxSize", 10),     // Max size in MB
		MaxBackups: GetIntOrDefault("log.file.maxBackups", 3),   // Max number of old log files to keep
		MaxAge:     GetIntOrDefault("log.file.maxAge", 28),      // Max age in days to keep a log file
		Compress:   GetBoolOrDefault("log.file.compress", true), // Compress old log files
	})
	logger.SetOutput(mw)
}

func getLogLevel() logrus.Level {
	lvl := strings.ToLower(GetStringOrDefault("log.level", "info"))
	switch lvl {
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
