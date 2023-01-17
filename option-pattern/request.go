package option_pattern

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// 针对可选的HTTP请求配置项，模仿gRPC使用的Options设计模式实现
type requestOption struct {
	timeout time.Duration
	data    string
	headers map[string]string
}

type Option struct {
	apply func(option *requestOption)
}

func defaultRequestOptions() *requestOption {
	return &requestOption{ // 默认请求选项
		timeout: 5 * time.Second,
		data:    "",
		headers: nil,
	}
}

func WithTimeout(timeout time.Duration) *Option {
	return &Option{
		apply: func(option *requestOption) {
			option.timeout = timeout
		},
	}
}

func WithData(data string) *Option {
	return &Option{
		apply: func(option *requestOption) {
			option.data = data
		},
	}
}

func httpRequest(method string, url string, options ...*Option) error {
	reqOpts := defaultRequestOptions() // 默认的请求选项
	for _, opt := range options {      // 在reqOpts上应用通过options设置的选项
		opt.apply(reqOpts)
	}
	// 创建请求对象
	req, err := http.NewRequest(method, url, strings.NewReader(reqOpts.data))
	if err != nil {
		return err
	}

	// 设置请求头
	for key, value := range reqOpts.headers {
		req.Header.Add(key, value)
	}
	// 发起请求
	client := &http.Client{Timeout: reqOpts.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println("responseCode:", resp.StatusCode)
	fmt.Println("print responseHead:")
	for key, value := range resp.Header {
		fmt.Println("\t\t", key, ":", value)
	}
	fmt.Println("print response body:")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	defer resp.Body.Close()
	return nil
}
