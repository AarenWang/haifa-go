package zaplog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func InitRollLogger() {
	encoder := getEncoder()

	hook := &lumberjack.Logger{
		Filename:   "./logs/roll_app" + ".log",
		MaxSize:    1,    //日志最大的大小（M）
		MaxBackups: 7,    //备份个数
		MaxAge:     7,    //最大保存天数（day）
		Compress:   true, //是否压缩
		LocalTime:  false,
	}

	w := zapcore.AddSync(hook)
	core := zapcore.NewCore(encoder, w, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller())
}

func Test_1(t *testing.T) {

	InitRollLogger()
	for {
		logger.Info("this is INFO level log", zap.String("name", "test"))
	}

}
