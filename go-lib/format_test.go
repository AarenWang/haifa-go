package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Teacher struct {
	Name    string
	ID      int `json:"id"`
	Age     int
	Address string `json:"address"`
}

var teacher = Teacher{
	Name:    "zhangsan",
	ID:      1,
	Age:     18,
	Address: "beijing",
}

func TestFormat(t *testing.T) {

	data, err := json.Marshal(teacher)

	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
		return
	}

	fmt.Printf("序列化之后的数据类型为: %T \n", data) // []uint8
	/**
	  查看序列化后的json字符串
	*/
	fmt.Println("序列化之后的数据为: ", string(data))

	tearcher2 := Teacher{}
	err2 := json.Unmarshal(data, &tearcher2)
	if err2 != nil {
		fmt.Println("反序列化出错,错误原因: ", err2)
		return
	}

	fmt.Printf("反序列化之后的数据为: %+v \n", tearcher2)
	fmt.Printf("反序列化之后的数据类型为: %T \n", tearcher2)
}

type AdvanceStruct struct {
	ID     int
	Name   string
	Sports []string
	Memory map[string]string
}

func TestFormat2(t *testing.T) {

	advanceStruct := AdvanceStruct{
		ID:     1,
		Name:   "zhangsan",
		Sports: []string{"basketball", "football"},
		Memory: map[string]string{"a": "apple", "b": "banana"},
	}

	data, err := json.Marshal(advanceStruct)

	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
		return
	}

	/**
	  查看序列化后的json字符串
	*/
	fmt.Println("序列化之后的数据为: ", string(data))

	advanceStruct2 := AdvanceStruct{}
	err2 := json.Unmarshal(data, &advanceStruct2)
	if err2 != nil {
		fmt.Println("反序列化出错,错误原因: ", err2)
		return
	}

	fmt.Printf("反序列化之后的数据为: %+v \n", advanceStruct2)
	fmt.Printf("反序列化之后的数据类型为: %T \n", advanceStruct2)
}
