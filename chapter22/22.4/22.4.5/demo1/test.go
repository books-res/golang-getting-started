package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 参数的原始内容
		rq := r.URL.RawQuery
		fmt.Printf("原始内容：%s\n", rq)
		// 处理后的参数列表
		fmt.Println("参数列表：")
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, strings.Join(v, ", "))
		}
		// 回复消息只包含头部，无正文
		// 状态码为200
		w.WriteHeader(http.StatusOK)
	})
	// 启动HTTP服务器
	http.ListenAndServe(":999", nil)
}
