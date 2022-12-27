package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/config/source/consul"
	"go-micro.dev/v4/config"
	"os"
)

func main() {
	load_from_consul()
	//load_from_file()
	//load_from_consul_json()
}

func load_from_consul() {
	//consul config
	configAddress := os.Getenv("CONFIG_ADDRESS")
	if configAddress == "" {
		configAddress = "127.0.0.1:8500"
	}

	consulSource := consul.NewSource(
		// optionally specify consul address; default to localhost:8500
		consul.WithAddress(configAddress),
		// 可选配置前缀为 /appconfig/entry
		consul.WithPrefix("appconfig/entry/"),
		// optionally strip the provided prefix from the keys, defaults to false
		consul.StripPrefix(true),
	)

	// Create new config
	conf, _ := config.NewConfig()
	// Load consul source
	if err := conf.Load(consulSource); err != nil {
		fmt.Println(err)
		return
	}

	tc := conf.Get("config.json", "topcity")
	fmt.Printf("topcity = %s \n", tc)

	fmt.Printf("get config item [topcity]: %s \n", conf.Get("config.json", "topcity"))
	var topCity []string
	//var topCity TopCity
	if err := conf.Get("config.json", "topcity").Scan(&topCity); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("topCity:", topCity)

}

type Host struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type TopCity struct {
	TopCity []string `json:"topcity"`
}
