package main

import (
	"craigslist.com/scraper/pkg"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

func setupLogs() *zap.Logger {
	const layout = "2006_01_02_15_04_05"
	t := time.Now().Format(layout)
	logFilePath, err := os.Getwd()
	logFileName := fmt.Sprintf("%s\\logs\\%s.log", logFilePath, t)
	logger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{logFileName},
		ErrorOutputPaths: []string{logFileName},
	}.Build()

	if err != nil {
		log.Fatalf("Failed to initialize zap logger: %v", err)
	}

	zap.ReplaceGlobals(logger)
	return logger
}

func main() {
	logger := setupLogs()
	defer logger.Sync()

	logger.Info("Initializing repository...")
	repo, err := pkg.NewRepository(logger)
	if err != nil {
		logger.Error("Failed to initialize repository", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("Fetching data...")
	searchList, err := repo.FetchData()
	if err != nil {
		logger.Error("Failed to fetch data", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("Fetching IDs...")
	vehicleList, err := repo.FetchIDs()
	if err != nil {
		logger.Error("Failed to fetch IDs", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("Initiating link check", zap.Int("TotalLinks", len(searchList)))

	var links []string

	for _, search := range searchList {
		logger.Info("Processing URL", zap.String("URL", search))
		searchLinks := pkg.PostListings(search, vehicleList, repo, logger)
		links = append(links, searchLinks...)
	}

	message := ""
	for _, link := range links {
		if len(message)+len(link) >= 1300 && len(message)+len(link) <= 1600 {
			pkg.SendMessage(message,logger)
			message = link + "\n"
		} else {
			message += link + "\n"
		}
	}

	if len(message) > 3 {
		logger.Info("Sending message...")
		pkg.SendMessage(message,logger)
		logger.Info("Raising alert...")
		pkg.Alert(logger)
	}

	logger.Info("Exiting application...")
}
