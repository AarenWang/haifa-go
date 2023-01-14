package concurrent

import (
	"log"
	"testing"
	"time"
)

func Test1(t *testing.T) {

	chn := make(chan int, 1)
	go func() {
		i := 1
		timer := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-timer.C:
				chn <- i
				log.Printf("write channel,value=%d \n", i)
				i++
			}
		}
	}()

	go func() {
		for {
			select {
			case d := <-chn:
				log.Printf("read channel,value=%d \n", d)

			}
		}

	}()

	select {}
}

type A struct {
	typeName string
}

type B struct {
	typeName string
}

func Test2(t *testing.T) {

	chn := make(chan interface{})
	go func() {
		i := 1
		timer := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-timer.C:
				a := A{typeName: "A"}
				chn <- a
				log.Printf("write channel,value=%v \n", a)
				i++
			}
		}
	}()

	go func() {
		for {
			select {
			case d := <-chn:
				log.Printf("read channel,value=%v \n", d)
			}
		}

	}()

	select {}
}
