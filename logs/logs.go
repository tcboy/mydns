package logs

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			"log_id": "@log_id",
		},
	})
	logger.SetOutput(os.Stdout)

}

func Info(format string, args ...interface{}) {
	logger.Infof(format, args)
}

func InfoCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Infof(format, args)
}

func Debug(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

func DebugCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Debugf(format, args)
}

func Warn(format string, args ...interface{}) {
	logger.Warnf(format, args)
}

func WarnCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Warnf(format, args)
}

func Error(format string, args ...interface{}) {
	logger.Errorf(format, args)
}

func ErrorCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Errorf(format, args)
}

func Fatal(format string, args ...interface{}) {
	logger.Fatalf(format, args)
}

func FatalCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Fatalf(format, args)
}

func Panic(format string, args ...interface{}) {
	logger.Panicf(format, args)
}

func PanicCtx(ctx context.Context, format string, args ...interface{}) {
	logger.WithContext(ctx).Panicf(format, args)
}
