package generic

import (
	"fmt"
	"testing"
)

// Generic is a generic type
func Test_1(t *testing.T) {
	type Slice[T int | float32 | float64] []T

	var s1 Slice[int] = []int{1, 2, 3}
	fmt.Printf("Type Name: %T \n", s1) // generic.Slice[int]

	var b Slice[float32] = []float32{1.0, 2.0, 3.0}
	fmt.Printf("Type Name: %T \n", b) // generic.Slice[float32]

	type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

	var m1 MyMap[int, float32] = map[int]float32{1: 1.0, 2: 2.0, 3: 3.0}
	fmt.Printf("Type Name: %T \n", m1) // generic.MyMap[int,float32]

}
