package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 待提交的数据
	var formdata = map[string][]string{
		"subject": {"some message"},
		"body":    {"the content of a message"},
		"order":   {"1"},
		"tags":    {"tag01", "tag02", "tag03"},
	}

	response, err := http.PostForm("http://localhost/upforms", formdata)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 输出服务器响应消息
	fmt.Print("服务器回应：")
	io.Copy(os.Stdout, response.Body)
}
