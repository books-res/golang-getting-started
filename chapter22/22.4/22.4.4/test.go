package main

import (
	"net/http"
)

func main() {
	// 设置Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("欢迎来到主页"))
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("“关于”本站"))
	})
	// 启动服务器
	http.ListenAndServe(":http", nil)
}
