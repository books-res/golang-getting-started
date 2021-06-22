package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		// 获取包含文件的字段
		if r.MultipartForm == nil {
			r.ParseMultipartForm(300000)
		}
		files, ok := r.MultipartForm.File["file"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if len(files) == 0 {
			w.Write([]byte("未上传任何文件"))
			return
		}
		// 保存文件
		var i int
		for _, fhd := range files {
			var fn = fhd.Filename
			fr, _ := fhd.Open()
			fileOut, _ := os.Create(fn)
			io.Copy(fileOut, fr)
			fileOut.Close()
			fr.Close()
			i++
		}
		msg := fmt.Sprintf("已成功上传%d个文件", i)
		w.Write([]byte(msg))
	})

	// 启动HTTP服务器
	http.ListenAndServe(":80", nil)
}
