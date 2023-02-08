package kv_serialize

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"log"
	"reflect"
	"strings"
	"time"
)

type KV[T any] struct {
	key   string
	value T
}

type KV_Store[T any] struct {
	kv         *KV[T]
	etcdClient clientv3.Client
}

func (s *KV_Store[T]) Store() {
	s.etcdClient.Put(context.Background(), s.kv.key, s.ValueToString())

}

func (s *KV_Store[T]) ValueToString() string {
	resType := reflect.TypeOf(s.kv.value)
	switch resType.Kind() {
	case reflect.Pointer:
		valType := resType.Elem()
		switch valType.Kind() {
		case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
			byte, err := json.Marshal(s.kv.value)
			if err != nil {
				panic(err)
			}
			return string(byte)
		case reflect.String:
			return reflect.ValueOf(s.kv.value).String()
		case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
			val := reflect.ValueOf(s.kv.value)
			derefVal := val.Elem()
			return fmt.Sprintf("%v", derefVal)
		}
		return ""
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		byte, err := json.Marshal(s.kv.value)
		if err != nil {
			panic(err)
		}
		return string(byte)
	case reflect.String:
		return reflect.ValueOf(s.kv.value).String()
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%v", s.kv.value)
	default:
		log.Printf("cannot resolve type %s \n", resType.Kind().String())
		return ""
	}
}

func (s *KV_Store[T]) Recovery(key string) T {
	getRsp, err := s.etcdClient.Get(context.Background(), key)
	if err != nil {
		panic(err)
	}

	rValue := parseBytes(getRsp.Kvs[0].Value, reflect.TypeOf(s.kv.value))

	return rValue.Interface().(T)
}

func (kv *KV[T]) Key() string {
	return kv.key
}

func parseBytes(b []byte, resType reflect.Type) reflect.Value {

	n := reflect.New(resType)
	if b == nil || len(b) == 0 {
		return n
	}
	switch resType.Kind() {
	case reflect.Pointer:
		child := parseBytes(b, resType.Elem())
		if !child.IsValid() {
			return reflect.Value{}
		}
		n.Elem().Set(child)
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		unmarshal := reflect.ValueOf(json.Unmarshal)
		args := []reflect.Value{
			reflect.ValueOf(b),
			n,
		}
		r := unmarshal.Call(args)
		if !r[0].IsNil() {
			return reflect.Value{}
		}
	case reflect.String:
		str := string(b)
		str = strings.Trim(str, "\"") //去掉双引号,否则 "test3" 会变成  "\"test3\""
		n.Elem().SetString(str)
	default:
		log.Printf("cannot resolve type %s \n", resType.Kind().String())
		return reflect.Value{}
	}
	return n
}

func Value[T any](kv *KV[T]) T {
	return kv.value
}

// Path: kv_serialize/kv_serialize.go

func EtcdClient() (*clientv3.Client, error) {
	clientCfg := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second, DialOptions: []grpc.DialOption{grpc.WithBlock()},
	}
	client, err := clientv3.New(clientCfg)
	if err != nil {
		return nil, err
	}

	return client, nil
}
