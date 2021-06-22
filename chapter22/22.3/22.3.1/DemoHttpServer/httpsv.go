/*
测试 HTTP 服务器的使用方法：
请先编译该代码文件
	go build -o <可执行文件> httpsv.go
然后运行生成的可执行文件

此测试服务器默认侦听80端口
测试时请以HTTP-GET方式访问 http://localhost/download

【注意】本测试程序中有一个sam.jpg文件，用于访问测试。
请将此文件与编译后的可执行文件放置在一起使用。
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var s = "欢迎访问【Demo Server】"
		w.Write([]byte(s))
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Write([]byte("请以GET方式访问"))
			return
		}
		http.ServeFile(w, r, "sam.jpg")
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Write([]byte("请以POST方式访问"))
			return
		}
		fmt.Println("客户端POST的数据：")
		reader := r.Body
		io.Copy(os.Stdout, reader)
		w.Write([]byte("提交成功"))
	})

	http.HandleFunc("/upforms", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Write([]byte("请以POST方式访问"))
			return
		}

		r.ParseForm()
		data := r.Form
		fmt.Println("客户端提交的数据：")
		for k, v := range data {
			fmt.Printf("\t%s: %s\n", k, strings.Join(v, ", "))
		}
		w.Write([]byte("提交完成"))
	})

	http.ListenAndServe(":http", nil)
}
