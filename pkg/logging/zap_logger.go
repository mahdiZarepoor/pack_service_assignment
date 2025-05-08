package logging

import (
	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var once sync.Once
var zapSinLogger *zap.SugaredLogger

type zapLogger struct {
	config configs.Config
	logger *zap.SugaredLogger
}

var zapLogLevelMapping = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func newZapLogger(config configs.Config) *zapLogger {
	logger := &zapLogger{config: config}
	logger.Init()
	return logger
}

func (l *zapLogger) getLogLevel() zapcore.Level {
	level, exists := zapLogLevelMapping[l.config.App.LogLevel]
	if !exists {
		return zapcore.DebugLevel
	}
	return level
}

func (l *zapLogger) Init() {
	once.Do(func() {
		stdoutWriter := zapcore.Lock(os.Stdout)

		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		stdoutCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			stdoutWriter,
			l.getLogLevel(),
		)

		logger := zap.New(stdoutCore, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
		zapSinLogger = logger.With("AppName", "MyApp").With("LoggerName", "ZapLog")
	})
	l.logger = zapSinLogger
}

func prepareLogKeys(cat Category, sub SubCategory, extra map[ExtraKey]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{})
	}
	extra["category"] = cat
	extra["subCategory"] = sub
	return mapToZapParams(extra)
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(cat, sub, extra)
	l.logger.Debugw(msg, params...)
}

func (l *zapLogger) DebugF(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(cat, sub, extra)
	l.logger.Infow(msg, params...)
}

func (l *zapLogger) InfoF(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(cat, sub, extra)
	l.logger.Warnw(msg, params...)
}

func (l *zapLogger) WarnF(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(cat, sub, extra)
	l.logger.Errorw(msg, params...)
}

func (l *zapLogger) ErrorF(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(cat, sub, extra)
	l.logger.Fatalw(msg, params...)
}

func (l *zapLogger) FatalF(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}
