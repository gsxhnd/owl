package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var logger *zap.Logger

var logLevel = zap.NewAtomicLevel()

func utcTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
}

func init() {
	//filePath := getFilePath()

	//w := zapcore.AddSync(&lumberjack.Logger{
	//	Filename:  filePath,
	//	MaxSize:   1024, //MB
	//	LocalTime: true,
	//	Compress:  true,
	//})

	//config := zap.NewProductionEncoderConfig()
	config := zap.NewDevelopmentEncoderConfig()
	//config.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	config.EncodeTime = utcTimeEncoder
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder

	core := zapcore.NewCore(
		//zapcore.NewJSONEncoder(config),
		zapcore.NewConsoleEncoder(config),
		os.Stdout,
		logLevel,
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

type Level int8

const (
	DebugLevel Level = iota - 1

	InfoLevel

	WarnLevel

	ErrorLevel

	DPanicLevel

	PanicLevel

	FatalLevel
)

func SetLevel(level Level) {
	logLevel.SetLevel(zapcore.Level(level))
}

//func getCurrentDirectory() string {
//	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
//	if err != nil {
//		log.Info(err)
//	}
//	return strings.Replace(dir, "\\", "/", -1)
//}
//
//func getFilePath() string {
//	logfile := getCurrentDirectory() + "/" + getAppname() + ".log"
//	return logfile
//}

//func getAppname() string {
//	full := os.Args[0]
//	full = strings.Replace(full, "\\", "/", -1)
//	splits := strings.Split(full, "/")
//	if len(splits) >= 1 {
//		name := splits[len(splits)-1]
//		name = strings.TrimSuffix(name, ".exe")
//		return name
//	}
//
//	return ""
//}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}
func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
