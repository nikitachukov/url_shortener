package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger = zap.NewNop()

func InitLogger() {
	var err error
	Log, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer Log.Sync()
}
