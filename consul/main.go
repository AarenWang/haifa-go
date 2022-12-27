package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	kv_optional(client, err)

	service_registery(client)

	//client.Agent().ServiceDeregister("myblog-1")

	service_query(client)

}

func service_query(client *api.Client) {
	services, _, err := client.Catalog().Services(nil)
	if err != nil {
		panic(err)
	}

	for name, tags := range services {
		fmt.Printf("service.name=%s, servicve.tag=%v \n", name, tags)
	}
}

func service_registery(client *api.Client) {
	blog := &api.AgentServiceRegistration{
		ID:      "myblog-1",
		Name:    "myblog",
		Tags:    []string{"primary"},
		Port:    9002,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			HTTP:     "http://127.0.0.1:1313/favicon.ico",
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	client.Agent().ServiceRegister(blog)

	blog2 := &api.AgentServiceRegistration{
		ID:      "myblog-2",
		Name:    "myblog",
		Tags:    []string{"primary"},
		Port:    9002,
		Address: "192.168.31.17",
		Check: &api.AgentServiceCheck{
			HTTP:     "http://192.168.31.17:1313/favicon.ico",
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	client.Agent().ServiceRegister(blog2)

}

func kv_optional(client *api.Client, err error) {
	kv := client.KV()

	p := &api.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	pair, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}
