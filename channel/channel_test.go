package channel

import (
	"log"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	boolChan := make(chan bool, 1)
	var currentTime string = "2020-01-01 00:00:00"

	go func() {
		defer close(boolChan)
		for {
			select {
			case <-boolChan:
				log.Printf("boolChan is ready,current time %s \n", currentTime)
			}
		}
	}()

	go func() {
		tick := time.Tick(1 * time.Second)

		for {
			select {
			case <-tick:
				boolChan <- true
				currentTime = time.Now().String()
			}
		}
	}()

	time.Sleep(100 * time.Second)
}
