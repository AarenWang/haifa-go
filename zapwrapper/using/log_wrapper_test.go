package using

import (
	"github.com/aarenwang/haifa-go/zapwrapper/wrapper/pkg/log"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func Test1(t *testing.T) {

	log.Debug("test debug")
	log.Info("test info")
	log.Warn("test warn")
	log.Error("test error")
	log.Fatal("test fatal")
	//log.Panic("test panic")
}

func Test2(t *testing.T) {
	logFile, err := os.Create("test.log")
	if err != nil {
		panic(err)
	}
	logger := log.New(logFile, log.InfoLevel)
	logger.Debug("test info")
	logger.Info("test info")

	logger.Info("test info", zap.String("key1", "value11111"), zap.Int("key2", 3333))

}

func Test3(t *testing.T) {

	loggerWritter := getRorateWritter()
	logger := log.New(loggerWritter, log.InfoLevel)
	logger.Debug("test debug") //debug级别的日志不会输出
	logger.Info("test info")

	logger.Info("test info", zap.String("key1", "value11111"), zap.Int("key2", 3333))
}

// 使用lumberjack实现日志切割
func getRorateWritter() zapcore.WriteSyncer {
	hook := lumberjack.Logger{
		Filename:   "./logs/app_rotate_log" + ".log",
		MaxSize:    1,    //日志最大的大小（M）
		MaxBackups: 10,   //备份个数
		MaxAge:     7,    //最大保存天数（day）
		Compress:   true, //是否压缩
		LocalTime:  false,
	}
	return zapcore.AddSync(&hook)
}
