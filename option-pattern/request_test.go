package option_pattern

import (
	"log"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {

	err := httpRequest("GET", "https://www.baidu.com/", WithTimeout(time.Second*10))
	if err != nil {
		t.Error(err)
	}

	log.Println("Hello")
}
