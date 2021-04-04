package log

import "go.uber.org/zap"

type ciLogger struct {
	logger *zap.Logger
}

func newCiLogger() *ciLogger {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	return &ciLogger{
		logger: logger,
	}
}

func (c *ciLogger) StartWait(message string) {

}

func (c *ciLogger) StopWait() {

}

func (c *ciLogger) Sync() {
	c.logger.Sync()
}

func (c *ciLogger) Debugf(format string, args ...interface{}) {
	zap.S().Debugf(format, args...)
}

func (c *ciLogger) Donef(format string, args ...interface{}) {
	zap.S().Infof(format, args...)
}

func (c *ciLogger) Infof(format string, args ...interface{}) {
	zap.S().Infof(format, args...)
}

func (c *ciLogger) Errorf(format string, args ...interface{}) {
	zap.S().Errorf(format, args...)
}

func (c *ciLogger) Warnf(format string, args ...interface{}) {
	zap.S().Warnf(format, args...)
}

func (c *ciLogger) Fatalf(format string, args ...interface{}) {
	zap.S().Fatalf(format, args...)
}
