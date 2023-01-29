package kratos_log

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"os"
	"testing"
)

func Test_1(t *testing.T) {
	logger := log.NewHelper(log.DefaultLogger) //std Log
	logger.Debug("This is Debug Level Log")
	logger.Info("This is Info Level Log")
	logger.Warn("This is Warn Level Log")
	logger.Error("This is Error Level Log")
}

// 设置LogLevel
func Test_2(t *testing.T) {
	logger := log.NewHelper(log.DefaultLogger, log.WithMessageKey("message"))

	logger.Debug("This is Debug Level Log")

	//调用Log(）方法,传入LogLevel key value
	logger.Log(log.LevelInfo, "key1", "value1")
	logger.Debug("This is Debug Level Log")
}

// Valuer 设置全局字段
func Test_3(t *testing.T) {
	logger := log.NewStdLogger(os.Stdout)
	logger = log.With(log.DefaultLogger, "caller", log.DefaultCaller, "ts", log.DefaultTimestamp)
	logger.Log(log.LevelInfo, "key1", "value1")
}

// Context:  logger.WithContext(ctx)
func Test_4(t *testing.T) {
	ctx := context.Background()
	logger := log.NewHelper(log.DefaultLogger)
	logger = logger.WithContext(ctx)
	logger.Debug("This is Debug Level Log")
}
