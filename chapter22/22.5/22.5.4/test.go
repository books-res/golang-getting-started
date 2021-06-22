package main

import (
	"net/http"
	"net/http/cgi"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 返回主页
		html := `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
</head>
<body>
    <form method="POST" action="/cmpt" enctype="multipart/form-data">
        <label for="file">请选择文件：</label>
        <input type="file" id="file" name="file" multiple />
        <input type="submit" value="上传" />
    </form>
</body>
</html>`
		w.Write([]byte(html))
	})

	// CGI程序
	var cgihandler = &cgi.Handler{
		Path: "test",
		Dir:  "./cgi-bin",
	}
	http.Handle("/cmpt", cgihandler)

	http.ListenAndServe(":http", nil)
}
