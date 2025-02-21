package logger

import (
	"os"

	"github.com/caophuoclong/whisper/configs"
	"github.com/caophuoclong/whisper/pkg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	cfg         *configs.Config
	sugarLogger *zap.SugaredLogger
}

func NewLogger(cfg *configs.Config) pkg.Logger {
	return &Logger{
		cfg: cfg,
	}
}

func (l *Logger) InitLogger() {
	// core := zapcore.
	logWrite := zapcore.AddSync(os.Stderr)
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"
	encoder := zapcore.NewConsoleEncoder(encoderCfg)
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWrite, zap.NewAtomicLevelAt(zapcore.InfoLevel))

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}
}

func (l *Logger) Debug(
	args ...interface{},
) {
}

func (l *Logger) Info(
	args ...interface{}) {
	l.sugarLogger.Info(args...)

}

func (l *Logger) Warn(
	args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *Logger) Error(
	args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *Logger) Fatal(
	args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}
