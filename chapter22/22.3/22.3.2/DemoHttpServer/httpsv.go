/*
测试HTTP服务器使用：
编译
	go build -o <可执行文件> httpsv.go
编译后可直接运行<可执行文件>
访问URL：http://localhost
*/

package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 读取头部信息
		fmt.Print("客户端发送的消息头：\n")
		for k, v := range r.Header {
			fmt.Printf("\t%s: %s\n", k, strings.Join(v, ", "))
		}
		// 向客户端发送回应消息
		rbmsg := "服务器已完成处理"
		w.Write([]byte(rbmsg))
	})
	http.ListenAndServe(":http", nil)
}
