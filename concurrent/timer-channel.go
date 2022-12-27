package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Timer-Channel Run")

	timer := time.NewTimer(2 * time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Printf("TimeOut current time is: %s \n", time.Now())
			timer.Reset(2 * time.Second)
		}

	}
}
