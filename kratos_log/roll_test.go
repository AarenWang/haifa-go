package kratos_log

import (
	"github.com/natefinch/lumberjack"
	"log"
	"testing"
)

func Test_roll_1(t *testing.T) {

	//timeLayout := time.RFC3339
	hook := lumberjack.Logger{
		Filename:   "./logs/applog" + ".log",
		MaxSize:    1,    //日志最大的大小（M）
		MaxBackups: 10,   //备份个数
		MaxAge:     7,    //最大保存天数（day）
		Compress:   true, //是否压缩
		LocalTime:  false,
	}

	log.SetOutput(&hook)

	for {
		log.Printf("This is a test log message")
	}

}
