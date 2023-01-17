package option_pattern

import (
	"testing"
	"time"
)

func TestRequest(t *testing.T) {

	err := httpRequest("GET", "https://www.baidu.com/", WithTimeout(time.Second*10))
	if err != nil {
		t.Error(err)
	}
}
