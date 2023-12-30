package logger

import (
	"log"
	"os"

	"github.com/aclgo/simple-api-gateway/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Start(*config.Config) error
	Info(args ...any)
	Infof(template string, args ...any)
	Warn(args ...any)
	Warnf(template string, args ...any)
	Error(args ...any)
	Errorf(template string, args ...any)
	Panic(args ...any)
	Panicf(template string, args ...any)
	Fatal(args ...any)
	Fatalf(template string, args ...any)
}

type apiLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func NewapiLogger(config *config.Config) (*apiLogger, error) {

	apl := &apiLogger{}

	if err := apl.Start(config); err != nil {
		return nil, err
	}

	return apl, nil
}

var (
	mapLogLevel = map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
		"panic": zapcore.PanicLevel,
		"fatal": zapcore.FatalLevel,
	}
)

func getLogLevel(cfg *config.Config) zapcore.Level {
	if level, ok := mapLogLevel[cfg.LogLevel]; ok {
		return level
	}

	return zapcore.DebugLevel
}

func (a *apiLogger) Start(cfg *config.Config) error {
	logLevel := getLogLevel(cfg)
	logWriter := zapcore.AddSync(os.Stderr)

	var encConfig zapcore.EncoderConfig

	if cfg.Server.Mode == "dev" {
		encConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encConfig = zap.NewProductionEncoderConfig()
	}

	encConfig.LevelKey = "LEVEL"
	encConfig.CallerKey = "CALLER"
	encConfig.TimeKey = "TIME"
	encConfig.NameKey = "NAME"
	encConfig.MessageKey = "MESSAGE"
	encConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder

	if cfg.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encConfig)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	a.sugaredLogger = logger.Sugar()

	if err := logger.Sync(); err != nil {
		log.Printf("logger.Sync: %v", err)
	}

	return nil
}

func (a *apiLogger) Info(args ...any) {
	a.sugaredLogger.Info(args...)
}
func (a *apiLogger) Infof(template string, args ...any) {
	a.sugaredLogger.Infof(template, args...)
}
func (a *apiLogger) Warn(args ...any) {
	a.sugaredLogger.Warn(args...)
}
func (a *apiLogger) Warnf(template string, args ...any) {
	a.sugaredLogger.Warnf(template, args...)
}
func (a *apiLogger) Error(args ...any) {
	a.sugaredLogger.Error(args...)
}
func (a *apiLogger) Errorf(template string, args ...any) {
	a.sugaredLogger.Errorf(template, args)
}
func (a *apiLogger) Panic(args ...any) {
	a.sugaredLogger.Panic(args...)
}
func (a *apiLogger) Panicf(template string, args ...any) {
	a.sugaredLogger.Panicf(template, args...)
}
func (a *apiLogger) Fatal(args ...any) {
	a.sugaredLogger.Fatal(args...)
}
func (a *apiLogger) Fatalf(template string, args ...any) {
	a.sugaredLogger.Fatalf(template, args...)
}
