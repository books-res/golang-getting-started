package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 创建请求
	var req, err = http.NewRequest(http.MethodGet, "http://localhost", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置自定义的HTTP头
	req.Header.Add("client-name", "Dabin")
	req.Header.Add("client-ver", "1.0")

	// 发起请求
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打印服务器响应消息
	fmt.Print("服务器响应：")
	io.Copy(os.Stdout, response.Body)
}
