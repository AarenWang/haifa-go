package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
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
	//go语言的bug：reflect_1.ValueOf(nil).IsNil()是不允许的，会报错
	//if !reflect_1.ValueOf(data).Elem().IsNil() {
	//	return
	//}

	n := reflect.New(reflect.TypeOf(data).Elem())
	//这样写编译不过
	//*data = n.Interface()
	v := reflect.ValueOf(data)
	v.SetPointer(unsafe.Pointer(n.UnsafeAddr()))
}

func TestFunc1(t *testing.T) {
	var test **Test

	//执行这个函数，使指针指向不为nil
	newElemWhenNil(test)

	err := json.Unmarshal(data, test)
	fmt.Println(test, err)

}

func TestFunc2(t *testing.T) {

	//panic: reflect_1: call of reflect_1.Value.IsNil on zero Value [recovered]
	//panic: reflect_1: call of reflect_1.Value.IsNil on zero Value
	data := func3()
	fmt.Printf("data Type: %T", data)

}

func func3() []byte {
	return nil
}

func InitPointerData(data any) {
	//确定为指针类型
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		return
	}
	//确定引用为nil
	if reflect.ValueOf(data).Elem().IsValid() {
		return
	}

	typeOf := reflect.TypeOf(data)
	// 显示反射类型对象的名称和种类
	fmt.Printf("typeOf name:'%v' kind:'%v'\n", typeOf.Name(), typeOf.Kind())
	typeOfElem := typeOf.Elem()
	fmt.Printf("typeOfElem name:'%v' kind:'%v'\n", typeOfElem.Name(), typeOfElem.Kind())

	//reflect_1.ValueOf(typeOf).SetPointer(unsafe.Pointer(reflect_1.ValueOf(typeOfElem).Pointer())) // panic: reflect_1: reflect_1.Value.SetPointer using unaddressable value
	//reflect_1.ValueOf(typeOfElem).Pointer() // uintptr

	//reflect_1.ValueOf(typeOf).SetPointer(reflect_1.ValueOf(typeOfElem).UnsafePointer()) //  panic: reflect_1: reflect_1.Value.SetPointer using unaddressable value [recovered]

	//va := reflect_1.ValueOf(data).Elem()
	////v := reflect_1.New(va.Type().Elem())
	//v := reflect_1.New(reflect_1.TypeOf(data).Elem())
	//va.SetPointer(unsafe.Pointer(v.UnsafeAddr()))
	data = reflect.New(typeOf)
}

func TestInitPointerData(t *testing.T) {
	var test *Test
	InitPointerData(test)
	fmt.Printf("test: %v", test)
}

func Initialize(v interface{}) {
	rv := reflect.ValueOf(v).Elem()
	rv.Set(reflect.New(rv.Type().Elem()))
}

func TestInitialize(t *testing.T) {
	var test *Test = &Test{}
	Initialize(test)
	fmt.Printf("test: %v", test)
}

func InitializePoint(v interface{}) {
	rv := reflect.ValueOf(v).Elem()
	rv.Set(reflect.New(rv.Type().Elem()))
}

func TestInitializePoint(t *testing.T) {
	var test *Test = &Test{}
	InitializePoint(test)
	fmt.Printf("test: %v", test)
}
