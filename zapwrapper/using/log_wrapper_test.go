package using

import (
	"github.com/aarenwang/haifa-go/zapwrapper/wrapper/pkg/log"
	"go.uber.org/zap"
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
