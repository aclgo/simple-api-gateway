package logger

import (
	"github.com/aclgo/simple-api-gateway/config"
	"go.uber.org/zap"
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

func (a *apiLogger) Start(*config.Config) error {
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
