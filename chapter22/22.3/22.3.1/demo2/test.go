package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	var strReader = strings.NewReader("Test Data")
	var response, err = http.Post("http://localhost/submit", "text/plain;charset=utf-8", strReader)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 接收服务器响应
	fmt.Print("服务器回应：")
	io.Copy(os.Stdout, response.Body)
}
