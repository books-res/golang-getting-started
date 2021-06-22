package main

import "net/http"

func main() {
	// 创建ServeMux实例
	svmux := http.NewServeMux()
	// 或者获取默认实例
	//svmux := http.DefaultServeMux
	// 建立URL路径与Handler的映射
	svmux.Handle("/pxo", new(myHandler))
	svmux.HandleFunc("/pxs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("测试页面-2"))
	})
	// 启动HTTP服务器
	var httpSvr = new(http.Server)
	httpSvr.Addr = "localhost:8080"
	httpSvr.Handler = svmux
	httpSvr.ListenAndServe()
}

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("测试页面-1"))
}
