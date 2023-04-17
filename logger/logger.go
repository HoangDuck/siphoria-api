package logger

import (
	"context"
	"fmt"
	"go.elastic.co/apm/module/apmzap"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"hotel-booking-api/logger/level"
	"hotel-booking-api/logger/rotate"
	"net/url"
	"os"
	"sync"
)

const (
	logFolder = "log"
	logFile   = "go.log"
)

var (
	logger  *zap.Logger
	syncOne sync.Once
)

func init() {
	_, err := os.Stat(logFolder)
	if os.IsNotExist(err) {
		if err := os.Mkdir(logFolder, os.ModePerm); err != nil {
			panic(err)
		}
	}

	if logger == nil {
		NewLogger(level.Debug)
	}
}

func Must(logger *zap.Logger, err error) *zap.Logger {
	if err != nil {
		panic(err)
	}
	return logger
}

func ViewLoggerDB(err *gorm.DB) {
	if err.Error != nil {
		logger.Error("SQL Error", zap.Error(err.Error))
		return
	}
	if err.RowsAffected == 0 {
		logger.Debug("SQL 0 row affected", zap.Error(err.Error))
		return
	}
	logger.Info("Info", zap.Error(err.Error))
}

func GetLogger() *zap.Logger {
	return logger
}

func NewLogger(lvl level.Level) {
	syncOne.Do(func() {
		if err := zap.RegisterSink("lumberjack", func(u *url.URL) (zap.Sink, error) {
			return rotate.LumberjackSink{
				Logger: rotate.Logger(logFile),
			}, nil
		}); err != nil {
			panic(err)
		}

		zap.NewProductionEncoderConfig()

		logger = Must(zap.Config{
			Level:    zap.NewAtomicLevelAt(level.ZapLevel(lvl)),
			Encoding: "json",
			EncoderConfig: ecszap.ECSCompatibleEncoderConfig(zapcore.EncoderConfig{
				EncodeLevel:  zapcore.CapitalLevelEncoder,
				EncodeCaller: zapcore.ShortCallerEncoder,
			}),
			OutputPaths:      []string{"stderr", fmt.Sprintf("lumberjack:%s", logFile)},
			ErrorOutputPaths: []string{"stderr"},
			Sampling:         nil,
		}.Build(zap.WrapCore((&apmzap.Core{}).WrapCore), zap.AddCaller()))

		zap.ReplaceGlobals(logger)
		zap.RedirectStdLog(logger)
	})
}

func TraceLogger(ctx context.Context) *zap.Logger {
	traceFields := apmzap.TraceContext(ctx)
	if v := ctx.Value("X-Request-ID"); v != nil {
		traceFields = append(traceFields, zap.String("ref_id", v.(string)))
	}
	return logger.With(traceFields...)
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	logger.Panic(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}

func Debugf(message string, args ...interface{}) {
	logger.Sugar().Debugf(message, args...)
}

func Infof(message string, args ...interface{}) {
	logger.Sugar().Infof(message, args...)
}

func Warnf(message string, args ...interface{}) {
	logger.Sugar().Warnf(message, args...)
}

func Errorf(message string, args ...interface{}) {
	logger.Sugar().Errorf(message, args...)
}

func Panicf(message string, args ...interface{}) {
	logger.Sugar().Panicf(message, args...)
}

func Fatalf(message string, args ...interface{}) {
	logger.Sugar().Fatalf(message, args...)
}
