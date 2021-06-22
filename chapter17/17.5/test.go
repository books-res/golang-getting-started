package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 创建新目录
	os.Mkdir("tmp", 0700)
	// 创建临时文件
	tmpfile, err := ioutil.TempFile("./tmp", "Demo*.tp")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("临时文件%s创建完成\n", tmpfile.Name())
	tmpfile.WriteString("abc")
	tmpfile.Close()
	fmt.Print("按任意键继续……")
	fmt.Scanln()
	// 删除文件
	os.Remove(tmpfile.Name())
}
