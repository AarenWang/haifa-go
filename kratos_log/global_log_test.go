package kratos_log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"

	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
)

func Test_g_1(t *testing.T) {
	//f, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	return
	//}
	writeSyncer := zapcore.AddSync(os.Stdout)

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	z := zap.New(core)

	logger := kratoszap.NewLogger(z)
	//logger = log.With(logger, "caller", log.DefaultCaller, "ts", log.DefaultTimestamp)
	log.SetLogger(logger)

	// 打印日志
	log.Info("info")
	log.Debug("debug")

	// 打印日志，带上字段

	zapLogger := log.NewHelper(logger)
	zapLogger.Infof("zapLogger info name=%s age=%s", "JackMa", "10000")

}

func TestValue(t *testing.T) {
	var DefaultLogger = log.NewStdLogger(os.Stdout)
	logger := DefaultLogger
	//logger 输出caller和时间戳
	logger = log.With(logger, "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)
	_ = logger.Log(log.LevelInfo, "msg", "hello world")

	// logger没有多余字段
	logger = DefaultLogger
	logger = log.With(logger)
	_ = logger.Log(log.LevelDebug, "msg", "hello world")

	var v1 interface{}
	got := log.Value(context.Background(), v1)
	if got != v1 {
		t.Errorf("Value() = %v, want %v", got, v1)
	}
	var v2 log.Valuer = func(ctx context.Context) interface{} {
		return 3
	}
	got = log.Value(context.Background(), v2)
	res := got.(int)
	if res != 3 {
		t.Errorf("Value() = %v, want %v", res, 3)
	}
}
