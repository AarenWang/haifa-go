package kv_serialize

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
)

func TestEtcdClient(t *testing.T) {
	client, err := EtcdClient()
	if err != nil {
		t.Fatal(err)
	}

	getRsp, err := client.Get(context.Background(), "test")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(getRsp)

}

func TestGet(t *testing.T) {
	client, err := EtcdClient()
	if err != nil {
		t.Fatal(err)
	}

	client.Put(context.Background(), "test/1", "1")
	client.Put(context.Background(), "test/2", "2")
	client.Put(context.Background(), "test/3", "3")
	getRsp, err := client.Get(context.Background(), "test", clientv3.WithPrefix())
	if err != nil {
		t.Fatal(err)
	}

	var kvs []*mvccpb.KeyValue = getRsp.Kvs
	for _, kv := range kvs {
		fmt.Println(string(kv.Key), string(kv.Value))
	}
}

type MyStruct struct {
	Name string
	Age  int
}

func TestKVStruct(t *testing.T) {
	kv_int := KV[int]{key: "test_int", value: 1}
	fmt.Println(kv_int.Key())

	kv_string := KV[string]{key: "test_string", value: "1"}
	fmt.Println(kv_string.Key())

	kv_map := KV[map[string]string]{key: "test_map", value: map[string]string{"1": "1"}}
	fmt.Println(kv_map.Key())

	kv_array := KV[[2]string]{key: "test_array", value: [2]string{"1", "2"}}
	fmt.Println(kv_array.Key())

	kv_struct := KV[MyStruct]{key: "test_struct", value: MyStruct{Name: "1", Age: 1}}
	fmt.Println(kv_struct.Key())

	kv_pointer := KV[*MyStruct]{key: "test_pointer", value: &MyStruct{Name: "1", Age: 1}}
	fmt.Println(kv_pointer.Key())

	kv_interface := KV[interface{}]{key: "test_interface", value: 1}
	fmt.Println(kv_interface.Key())

}

func TestValueToString(t *testing.T) {
	mapStore := KvStore[map[string]string]{
		kv: &KV[map[string]string]{key: "test_map", value: map[string]string{"1": "1"}},
	}
	strVal := mapStore.ValueToString()
	fmt.Println("map_store:", strVal)

	arrayStore := KvStore[[2]string]{
		kv: &KV[[2]string]{key: "test_array", value: [2]string{"1", "2"}},
	}
	strVal = arrayStore.ValueToString()
	fmt.Println("array_store:", strVal)

	pointerStore := KvStore[*MyStruct]{
		kv: &KV[*MyStruct]{key: "test_pointer", value: &MyStruct{Name: "1", Age: 1}},
	}
	strVal = pointerStore.ValueToString()
	fmt.Println("pointer_store:", strVal)

	stringStore := KvStore[string]{
		kv: &KV[string]{key: "test_string", value: "1"},
	}

	strVal = stringStore.ValueToString()
	fmt.Println("string_store:", strVal)

	intStore := KvStore[int]{
		kv: &KV[int]{key: "test_int", value: 1},
	}

	strVal = intStore.ValueToString()
	fmt.Println("int_store:", strVal)

	float64Store := KvStore[float64]{
		kv: &KV[float64]{key: "test_int", value: float64(10.01)},
	}
	strVal = float64Store.ValueToString()
	fmt.Println("float64_store:", strVal)

	var i int = 10
	intPointerStore := KvStore[*int]{
		kv: &KV[*int]{key: "test_pointer_int", value: &i},
	}
	strVal = intPointerStore.ValueToString()
	fmt.Println("int_pointer_store:", strVal)
}

func TestKvStore_Recovery(t *testing.T) {
	client, err := EtcdClient()
	if err != nil {
		panic("")
	}

	stringStore := KvStore[string]{
		kv:         &KV[string]{key: "test_string", value: "1"},
		etcdClient: client,
	}

	strVal := stringStore.ValueToString()
	fmt.Println("string_store:", strVal)

	arrayStore := KvStore[[2]string]{
		kv:         &KV[[2]string]{key: "test_array", value: [2]string{"1", "2"}},
		etcdClient: client,
	}
	strVal = arrayStore.ValueToString()
	fmt.Println("array_store:", strVal)

	mapStore := KvStore[map[string]string]{
		kv:         &KV[map[string]string]{key: "test_map", value: map[string]string{"1": "1"}},
		etcdClient: client,
	}
	strVal = mapStore.ValueToString()
	fmt.Println("map_store:", strVal)
	mapStore.Store()
	mapStore.Recovery()

	pointerStore := KvStore[*MyStruct]{
		kv:         &KV[*MyStruct]{key: "test_pointer", value: &MyStruct{Name: "1", Age: 1}},
		etcdClient: client,
	}
	strVal = pointerStore.ValueToString()
	fmt.Println("pointer_store:", strVal)
	pointerStore.Store()
	pointerStore.Recovery()

	intStore := KvStore[int]{
		kv:         &KV[int]{key: "test_int", value: 1},
		etcdClient: client,
	}

	strVal = intStore.ValueToString()
	fmt.Println("int_store:", strVal)

	float64Store := KvStore[float64]{
		kv:         &KV[float64]{key: "test_int", value: float64(10.01)},
		etcdClient: client,
	}
	strVal = float64Store.ValueToString()
	fmt.Println("float64_store:", strVal)

	var i int = 10
	intPointerStore := KvStore[*int]{
		kv:         &KV[*int]{key: "test_pointer_int", value: &i},
		etcdClient: client,
	}
	strVal = intPointerStore.ValueToString()
	fmt.Println("int_pointer_store:", strVal)
}
