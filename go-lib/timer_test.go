package main

import (
	"fmt"
	"testing"
	"time"
)

// 定时器 使用方法1
func Test1(t *testing.T) {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("timer1: Time out!")

	// After 直接返回通道  C <-chan Time
	timer2Chan := time.After(time.Second * 2)
	<-timer2Chan
	fmt.Println("timer: Time out!")

	// After 返回的通道 可以用select 语句
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("timer3: Time out!")
	}

	//取消定时器  Stop()  返回true 代表成功取消，返回false 代表定时器已经超时或者被取消
	timer4 := time.NewTimer(time.Second * 2)
	if timer4.Stop() {
		fmt.Println("timer4: Stop!")
	}
}

// 周期定时器
func Test2(t *testing.T) {
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for {
			<-ticker.C
			fmt.Println("ticker: Time Reach!")
		}
	}()

	//等待周期定时器 执行5次
	select {
	case <-time.After(time.Second * 10):
		fmt.Println("timer: Time out!")
	}

}
