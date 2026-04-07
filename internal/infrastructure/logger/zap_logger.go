package logger

import (
	"fmt"
	"go-simple-api/internal/config"
	"go-simple-api/internal/utils"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func ZapLogger(cfg config.Config) (*zap.Logger, error) {

	filename := fmt.Sprintf("logs/%s.log", cfg.Logger.LogFilename)

	logWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    utils.StringToInt(cfg.Logger.LogMaxSize),    // MB
		MaxBackups: utils.StringToInt(cfg.Logger.LogMaxBackups), // file
		MaxAge:     utils.StringToInt(cfg.Logger.LogMaxAge),     // days
		Compress:   utils.StringToBool(cfg.Logger.LogCompress),
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
