package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/http/cgi"
	"path/filepath"
	"strings"
)

func main() {
	// 获取Request对象
	var req, err = cgi.Request()
	if err != nil {
		// 响应错误消息
		fmt.Print("HTTP/1.1 500 服务器内部发生了错误\r\n")
		return
	}
	if req.Method != http.MethodPost {
		fmt.Print("HTTP/1.1 405 请以POST方式提交\r\n")
		return
	}

	// 读取文件
	req.ParseMultipartForm(7988850000)
	files := req.MultipartForm.File["file"]
	strbd := new(strings.Builder)
	for _, f := range files {
		// 获取短文件名
		var fn = filepath.Base(f.Filename)
		// 获取文件内容
		thefile, _ := f.Open()
		reader := io.LimitReader(thefile, f.Size)
		// 计算MD5值
		md5 := md5.New()
		io.Copy(md5, reader)
		var md5data = md5.Sum(nil)
		// 生成文本
		h := fmt.Sprintf(`<p>文件%s的MD5：%x</p>`, fn, md5data)
		strbd.WriteString(h + "\r\n")
	}

	// 写入响应消息
	// 写入状态行
	// 此行包括：
	// 1、HTTP版本号为1.1
	// 2、状态码为200
	// 3、状态码文本为“OK”
	fmt.Print("HTTP/1.1 200 OK\r\n")
	// 写入消息头
	fmt.Print("Content-Type: text/html;charset=utf-8\r\n")
	// 写入分隔行（空行）
	fmt.Print("\r\n")
	// 构建<body>内容
	var body string
	if strbd.Len() == 0 {
		body = `<body>
			<p>未发现任何文件</p>
		</body>`
	} else {
		body = fmt.Sprintf(`<body>
		<h3>处理结果</h3>
			%s
		</body>`, strbd)
	}
	// 写入HTML文档
	fmt.Print("<!DOCTYPE html>\r\n")
	fmt.Printf(`<html>
	%s
</html>`, body)
}
