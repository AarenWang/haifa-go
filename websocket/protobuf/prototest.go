package main

import (
	"fmt"
	"github.com/aarenwang/go-haifa/websocket/protobuf/api"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func main1() {

	client1 := &api.Client{
		RemoteAddr:  "192.168.3.5",
		LocalAddr:   "192.168.3.4",
		SubProtocol: "my-protocol-v1",
	}

	client2 := &api.Client{
		RemoteAddr:  "192.168.3.6",
		LocalAddr:   "192.168.3.4",
		SubProtocol: "my-protocol-v2",
	}

	clientList := &api.ClientList{
		Clients: []*api.Client{client1, client2},
	}

	out, err := proto.Marshal(clientList)
	if err != nil {
		fmt.Printf("proto Marshal err %s \n", err)
		return
	}
	if err := ioutil.WriteFile("client.proto", out, 0644); err != nil {
		fmt.Printf("write to client.proto fail err=%s \n", err)
	}

	in, err2 := ioutil.ReadFile("client.proto")
	if err2 != nil {
		fmt.Printf("reate  client.proto fail err=%s \n", err2)
	}

	clientList2 := &api.ClientList{}
	if err := proto.Unmarshal(in, clientList2); err != nil {
		fmt.Printf("proto Unmarshal err %s \n", err)
		return
	}
	fmt.Printf("clientList= %v \n", clientList2)
}
