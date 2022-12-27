package mytest

import (
	"fmt"
	"log"
	"testing"
)

func Test_panic_and_recover(t *testing.T) {

	log.Println("Test_panic_and_recover")

	red := Apple{color: "red"}
	green := Apple{color: "green"}
	r1 := check_apple_color(red)
	r2 := check_apple_color(green)
	log.Printf("r1=%s \n ", r1)
	log.Printf("r2=%s \n ", r2)
}

func check_apple_color(apple Apple) string {

	defer func() {
		//fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}

	}()

	if apple.color == "green" {
		panic("Oh my God green apple")
	}
	return apple.color
}

func test() (a int) {
	defer func() {
		if err := recover(); err != nil {
			a = 1
		}
	}()
	var p *int
	*p = 0
	return 0
}
