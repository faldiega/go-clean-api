package config

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() (*zap.Logger, error) {

	date := time.Now().Format("2006-01-02")

	filename := fmt.Sprintf("logs/app_%s.log", date)

	logWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // MB
		MaxBackups: 5,  // file
		MaxAge:     1,  // days
		Compress:   true,
	})

	consoleWriter := zapcore.AddSync(os.Stdout)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, logWriter, zap.InfoLevel),
		zapcore.NewCore(consoleEncoder, consoleWriter, zap.InfoLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return logger, nil
}
