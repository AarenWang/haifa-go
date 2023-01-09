package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	colors0 := make([]string, 0) // colors0 length: 0, capacity: 0
	fmt.Printf("colors0 length: %d, capacity: %d\n", len(colors0), cap(colors0))

	colors1 := make([]string, 5)
	colors1[0] = "red"

	colors2 := make([]string, 5, 10)
	colors2[0] = "red"

	colors3 := make([]string, 0, 10)
	//colors3[0] = "red" // 容量为10,长度为0，不能赋值,需要通过append
	colors3 = append(colors3, "red")

	colors4 := []string{"red", "blue", "green", "yellow"} // length: 4, capacity: 4
	fmt.Printf("colors5 length: %d, capacity: %d\n", len(colors4), cap(colors4))

}
