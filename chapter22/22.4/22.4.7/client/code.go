package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buffer = new(bytes.Buffer)
	multWrt := multipart.NewWriter(buffer)
	// 生成三个文件字段
	fw, _ := multWrt.CreateFormFile("file", "file-1.txt")
	fw.Write([]byte("测试文件 #1"))
	fw, _ = multWrt.CreateFormFile("file", "file-2.txt")
	fw.Write([]byte("测试文件 #2"))
	fw, _ = multWrt.CreateFormFile("file", "file-3.txt")
	fw.Write([]byte("测试文件 #3"))
	// 关闭Writer
	multWrt.Close()

	// 创建Request对象
	reader := bytes.NewReader(buffer.Bytes())
	var req, err = http.NewRequest(http.MethodPost, "http://127.0.0.1/upload", reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置Content-Type头
	req.Header.Set("Content-Type", multWrt.FormDataContentType())
	// 发送请求
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("服务器回应：")
	io.Copy(os.Stdout, response.Body)
}
