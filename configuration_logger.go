package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/url"
	"os"
	"strings"
)

var log *zap.SugaredLogger

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}
func initLogger() {
	enableFileLogging := GetBoolOrDefault("log.file.enable", true)
	logFolder := viper.GetString("log.folder")
	outputPaths := []string{"stdout"}
	errorPaths := []string{"stderr"}

	if enableFileLogging {
		err := os.MkdirAll(logFolder, os.ModePerm)
		if err != nil {
			fmt.Errorf("error creating log folder: %v", err)
			os.Exit(1)
		}

		ll := lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s", logFolder, "configuration.log"),
			MaxSize:    GetIntOrDefault("log.file.maxSize", 10),     // Max size in MB
			MaxBackups: GetIntOrDefault("log.file.maxBackups", 3),   // Max number of old log files to keep
			MaxAge:     GetIntOrDefault("log.file.maxAge", 28),      // Max age in days to keep a log file
			Compress:   GetBoolOrDefault("log.file.compress", true), // Compress old log files
		}
		err = zap.RegisterSink("config", func(*url.URL) (zap.Sink, error) {
			return lumberjackSink{
				Logger: &ll,
			}, nil
		})
		if err != nil {
			panic(fmt.Sprintf("build zap logger from config error: %v", err))
		}
		outputPaths = append(outputPaths, fmt.Sprintf("config:%s", fmt.Sprintf("%s/%s", logFolder, fmt.Sprintf("%s/%s", logFolder, "configuration.log"))))
	}

	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(getLogLevel()),
		OutputPaths:      outputPaths,
		ErrorOutputPaths: errorPaths,
		Development:      GetBoolOrDefault("log.dev", false),
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:      "level",
			TimeKey:       "time",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			CallerKey:     "caller",
			EncodeCaller:  zapcore.ShortCallerEncoder,
			NameKey:       "logger",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			//EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
		},
	}
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(fmt.Sprintf("build zap logger from config error: %v", err))
	}
	defer logger.Sync() // flushes buffer, if any
	log = logger.Sugar()
}

func getLogLevel() zapcore.Level {
	lvl := strings.ToLower(GetStringOrDefault("log.level", "info"))
	switch lvl {
	case "info":
		return zapcore.InfoLevel
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel

	}
}
