package consul

import (
	"fmt"
	"testing"
)

func TestInitClient(t *testing.T) {
	//InitClient()
	client := InitClient()
	fmt.Print(client.Catalog().Services(nil))
}

func TestService_query(t *testing.T) {
	client := InitClient()
	ServiceQuery(client)
}

func TestServiceRegistry(t *testing.T) {
	client := InitClient()
	ServiceRegistry(client)
}

func TestAgentServices(t *testing.T) {
	client := InitClient()
	AgentServices(client)
}

func TestServiceDeregister(t *testing.T) {
	client := InitClient()
	ServiceDeregister(client, "myblog-3")
}
