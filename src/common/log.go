package common

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func init() {
	noSugarLogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	Logger = noSugarLogger.Sugar()
	defer Logger.Sync()
}
