package logger

import "go.uber.org/zap"

func New() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()

	sugar.Infow("failed to fetch url", "url", "http://google.com")

	return sugar
}
