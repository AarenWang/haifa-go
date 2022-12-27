package main

import (
	"fmt"
	"time"
)

func main() {
	//当前时间
	now := time.Now()

	//时间转时间戳，单位秒
	seconds_stamp := now.Unix()
	//时间转字符串，时间格式化，模板为2006-01-02 15:04:05，区别为java中yyyy-MM-dd HH:mm:ss
	//解析在方法time.format.nextStdChunk()，用上符号表示年月日等，是因为各字段是互斥的，具体可以看源码
	fmt.Println(now.Format("2006-01-02 15:04:05")) //2021-07-24 10:30:52

	//时间戳转时间
	fmt.Println(time.Unix(seconds_stamp, 0)) // 输出 2021-07-24 10:30:52 +0800 CST
	//字符串转时间
	fmt.Println(time.Parse("2006-01-02 15:04:05", "2019-07-01 08:18:18"))
}
