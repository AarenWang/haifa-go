package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url = "https://blockscout.com/xdai/mainnet/tx/0x6d7c958517f5892a0b0fd7b00aeeb62c0876a0a6978cef1e4b3b6a4a8f125653"
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
