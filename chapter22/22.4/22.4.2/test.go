package main

import (
	"errors"
	"net/http"
	"strings"
)

func main() {
	// 构建服务器组件
	sv := new(http.Server)
	// 设置侦听地址
	sv.Addr = ":88"
	// 设置URL路径与Handler的映射
	var root = new(HandlerRoot)
	root.MapHandler("/", &Handler1{})
	root.MapHandler("/photos", &Handler2{})
	root.MapHandler("/musics", &Handler3{})
	sv.Handler = root
	// 启动HTTP服务器
	sv.ListenAndServe()
}

type HandlerRoot struct {
	// 该字段是映射类型，用于存放URL路径与Handler类型之间的对应关系
	hm map[string]http.Handler
}

func (h *HandlerRoot) MapHandler(path string, handler http.Handler) error {
	if path == "" {
		return errors.New("指定的URL不能为空")
	}
	// 让URL路径以“/”开头
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	// 如果目标已存在，不必再添加
	if _, isExist := h.hm[path]; isExist {
		return errors.New("此URL路径已经存在")
	}
	// hm字段有可能为nil
	if h.hm == nil {
		h.hm = make(map[string]http.Handler)
	}
	h.hm[path] = handler
	return nil
}

func (h *HandlerRoot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var urlpath = r.URL.Path
	hr, exist := h.hm[urlpath]
	if !exist {
		// 找不到匹配的路径
		hr = http.NotFoundHandler()
	}
	// 调用目标Handler
	hr.ServeHTTP(w, r)
}

/****************************************/
// 映射到URL路径“/”
type Handler1 struct{}

func (h *Handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "欢迎来到【我的小站】"
	w.Write([]byte(msg))
}

// 映射到URL路径“/photos”
type Handler2 struct{}

func (h *Handler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "欢迎使用【我的相册】"
	w.Write([]byte(msg))
}

// 映射到URL路径“/musics”
type Handler3 struct{}

func (h *Handler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "欢迎使用【我的音乐】"
	w.Write([]byte(msg))
}
