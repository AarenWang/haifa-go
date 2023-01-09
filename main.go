package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

type Test struct {
	A string `json:"a"`
	B string `json:"b"`
}

var data = []byte("{\"a\":\"a\",\"b\":\"b\"}")

// 入参类型不能固定
func newElemWhenNil(data any) {
	//确定为指针类型
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		return
	}
	//确定引用为nil
	if reflect.ValueOf(data).Elem().IsValid() {
		return
	}
	//不知要如何判断是否nil
	//go语言的bug：reflect.ValueOf(nil).IsNil()是不允许的，会报错
	//if !reflect.ValueOf(data).Elem().IsNil() {
	//	return
	//}

	n := reflect.New(reflect.TypeOf(data).Elem())
	//这样写编译不过
	//*data = n.Interface()
	v := reflect.ValueOf(data)
	v.SetPointer(unsafe.Pointer(n.UnsafeAddr()))
}

func main() {
	var test *Test

	//执行这个函数，使指针指向不为nil
	newElemWhenNil(test)

	err := json.Unmarshal(data, test)
	fmt.Println(test, err)
}
