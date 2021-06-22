package main

import (
	"net/http"
	"net/http/cgi"
)

func main() {
	var (
		hd1 = &cgi.Handler{
			Dir:  "./cgi-bin",
			Path: "test1",
		}
		hd2 = &cgi.Handler{
			Dir:  "./cgi-bin",
			Path: "test2",
		}
	)
	// 注册Handler
	http.Handle("/demo1", hd1)
	http.Handle("/demo2", hd2)
	// 启动HTTP服务器
	http.ListenAndServe(":http", nil)
}
