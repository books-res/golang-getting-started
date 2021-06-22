package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var response, err = http.Get("http://localhost/download")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 读取服务器回应的文件内容
	// 并保存到本地文件中
	outfile, err := os.Create("123.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outfile.Close()
	// 直接复制流中数据
	io.Copy(outfile, response.Body)
	fmt.Println("文件下载完成")
}
