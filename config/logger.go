package config

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger() {
	var err error

	env := os.Getenv("APP_ENV")

	if env == "production" {
		// 1) tentukan base path (aman untuk VPS / shared hosting)
		baseDir, _ := os.Getwd()
		logDir := filepath.Join(baseDir, "logs")

		// 2) pastikan folder logs ada
		if err := os.MkdirAll(logDir, 0755); err != nil {
			panic(err)
		}

		// 3) absolute path (ini kunci biar tidak nyasar)
		appLogPath := filepath.Join(logDir, "app.log")
		errLogPath := filepath.Join(logDir, "error.log")

		cfg := zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
			Development: false,
			Encoding:    "json",
			OutputPaths: []string{
				"stdout",
				appLogPath,
			},
			ErrorOutputPaths: []string{
				"stderr",
				errLogPath,
			},
			EncoderConfig: zap.NewProductionEncoderConfig(),
		}

		Log, err = cfg.Build()
		if err != nil {
			panic(err)
		}

		// 4) test write (biar pasti file dibuat)
		Log.Info("logger initialized",
			zap.String("env", env),
			zap.String("app_log", appLogPath),
			zap.String("error_log", errLogPath),
		)

	} else {
		Log, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
	}
}
