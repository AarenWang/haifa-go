package reflect_1

import (
	"encoding/json"
	"fmt"
	"github.com/bmizerany/assert"
	"reflect"
	"testing"
)

func TestF1(t *testing.T) {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}

// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

func TestF2(t *testing.T) {

	// 声明一个空结构体
	type cat struct {
	}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	// 获取Zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	json.Unmarshal([]byte{}, nil)
}

func Initialize(v interface{}) {
	//确定为指针类型
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return
	}

	// 再判断下  指针是否初始化
	if !IsValid(v) {
		return
	}

	rv := reflect.ValueOf(v).Elem()
	rv.Set(reflect.New(rv.Type().Elem()))
}

func IsValid(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsValid()
	}
	return false
}

type MyType struct {
	Name string
}

func TestF3(t *testing.T) {
	var val *MyType
	Initialize(&val)
	fmt.Println(val != nil)
}

type CType struct {
	A string `json:"a"`
	B string `json:"b"`
}

func TestF4(t *testing.T) {
	var ctype *CType
	fmt.Printf("before Init ctype address = %p\n", ctype)
	Initialize(&ctype)
	fmt.Printf("after Init ctype address = %p\n", ctype)
	Initialize(&ctype)
	fmt.Printf("second Init ctype address = %p\n", ctype)

	fmt.Println(ctype != nil)

	json.Unmarshal([]byte("{\"a\":\"a\",\"b\":\"b\"}"), ctype)
	fmt.Println(ctype)

	assert.Equal(t, ctype.A, "a")
	assert.Equal(t, ctype.B, "b")
}
