package logs

import (
	"github.com/gogotsenghsien/simple-rate-limit/src/configs"
	"github.com/sirupsen/logrus"
)

const (
	FieldError = "err"
	FieldIP    = "ip"
)

type Logger struct {
	*logrus.Entry
}

func NewLogger(config *configs.Config) (*Logger, error) {
	newLogger := logrus.New()
	threshold := config.GetString("logs.threshold")
	level, _ := logrus.ParseLevel(threshold)
	newLogger.SetLevel(level)
	newLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	})
	return &Logger{newLogger.WithField("logger", "default")}, nil
}

func (l Logger) RearrangeFields(pairs ...interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	// pairs must be even
	if len(pairs)%2 != 0 {
		return nil
	}

	// transfer pairs to map
	for i := 0; i < len(pairs); i += 2 {
		if key, ok := pairs[i].(string); ok {
			value := pairs[i+1]
			result[key] = value
		}
	}
	return result
}

func (l Logger) Debug(msg string, pairs ...interface{}) {
	fields := l.RearrangeFields(pairs...)
	l.WithFields(fields).Debugln(msg)
}

func (l Logger) Info(msg string, pairs ...interface{}) {
	fields := l.RearrangeFields(pairs...)
	l.WithFields(fields).Infoln(msg)
}

func (l Logger) Warn(msg string, pairs ...interface{}) {
	fields := l.RearrangeFields(pairs...)
	l.WithFields(fields).Warnln(msg)
}

func (l Logger) Error(msg string, pairs ...interface{}) {
	fields := l.RearrangeFields(pairs...)
	l.WithFields(fields).Errorln(msg)
}

func (l Logger) Fatal(msg string, pairs ...interface{}) {
	fields := l.RearrangeFields(pairs...)
	l.WithFields(fields).Fatalln(msg)
}

func (l Logger) Panic(msg string, pairs ...interface{}) {
	fields := l.RearrangeFields(pairs...)
	l.WithFields(fields).Panicln(msg)
}
