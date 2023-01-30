package zaplog

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"testing"
	"time"
)

// 使用Sugar Logger
func TestSugarLogger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "https://example.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

// 默认生产环境配置
func TestDefaultConfig(t *testing.T) {
	//logger, err := zap.NewProduction()
	logger, err := zap.NewDevelopment()
	zap.NewExample()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("This is a production config log example", zap.String("username", "JackMa"))

	// Production
	// {"level":"info","ts":1674984972.10138,"caller":"zaplog/zap_test.go:36","msg":"This is a production config log example","username":"JackMa"}

	// Development
	// 2023-01-29T21:41:03.362+0800	INFO	zaplog/zap_test.go:37	This is a production config log example	{"username": "JackMa"}
}

func CustomerConfigExample(options ...zap.Option) *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder, //日志级别字段带颜色，Encoder为ConsoleEncoder才有效
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zap.DebugLevel)
	return zap.New(core).WithOptions(options...)
}

func TestCustomerConfigExample(t *testing.T) {
	logger := CustomerConfigExample()
	logger.Info("This is a Debug Level customer config log example", zap.String("username", "JackMa"))
	logger.Info("This is a Info Level customer config log example", zap.String("username", "JackMa"))
	logger.Error("This is a Error Level customer config log example", zap.String("username", "JackMa"))

	// {"level":"info","ts":1674984972.10138,"caller":"zaplog/zap_test.go:36","msg":"This is a production config log example","username":"JackMa"}

}

func TestLogger(t *testing.T) {
	// logger, _ := zap.NewProduction()
	logger, _ := zap.NewDevelopment()

	defer logger.Sync()
	url := "https://example.com"
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

var logger *zap.Logger

// zapcore.NewCore 由三个参数组成，分别是编码器、写入器、日志级别
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller())
}

func TestInitLogger(t *testing.T) {
	InitLogger()

	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	logger.DPanic("dpanic")

	fmt.Println("done")
}

func TestRotateLogger(t *testing.T) {

	rorateWritter := getRorateWritter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, rorateWritter, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
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
