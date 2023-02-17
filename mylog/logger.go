package mylog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(level string) zapcore.Level {
	if level, ok := levelMap[level]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func init() {
	fileName := "zap.log"
	level := getLoggerLevel("debug")

	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})

	encoder := zap.NewProductionEncoderConfig()

	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewTee(
		// zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, level),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncWriter, level),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), zapcore.AddSync(os.Stdout), level),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))

	Logger = logger.Sugar()

}

// func Debug(args ...interface{}) {
// 	sugaredlogger.Debug(args...)
// }

// func Debugf(templ string, args ...interface{}) {
// 	sugaredlogger.Debugf(templ, args...)
// }

// func Info(args ...interface{}) {
// 	sugaredlogger.Info(args...)
// }

// func Infof(templ string, args ...interface{}) {
// 	sugaredlogger.Infof(templ, args...)
// }

// func Warn(args ...interface{}) {
// 	sugaredlogger.Warn(args...)
// }

// func Warnf(templ string, args ...interface{}) {
// 	sugaredlogger.Warnf(templ, args...)
// }

// func Error(args ...interface{}) {
// 	sugaredlogger.Error(args...)
// }

// func Errorf(templ string, args ...interface{}) {
// 	sugaredlogger.Errorf(templ, args...)
// }

// func DPanic(args ...interface{}) {
// 	sugaredlogger.DPanic(args...)
// }

// func DPanicf(templ string, args ...interface{}) {
// 	sugaredlogger.DPanicf(templ, args...)
// }

// func Panic(args ...interface{}) {
// 	sugaredlogger.Panic(args...)
// }

// func Panicf(templ string, args ...interface{}) {
// 	sugaredlogger.Panicf(templ, args...)
// }

// func Fatal(args ...interface{}) {
// 	sugaredlogger.Fatal(args...)
// }

// func Fatalf(templ string, args ...interface{}) {
// 	sugaredlogger.Fatalf(templ, args...)
// }
