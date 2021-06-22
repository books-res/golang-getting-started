package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"os"
)

func main() {
	var fileName = "demo.bin"
	// 创建文件并写入内容
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	var buffer = make([]byte, 3000)
	// 产生随机字节序列
	rand.Read(buffer)
	file.Write(buffer)
	file.Close()

	// 计算文件的SHA1值
	sha1 := sha1.New()
	inputFile, _ := os.Open(fileName)
	// 把文件的内容写入到Hash对象的缓冲区中
	io.Copy(sha1, inputFile)
	inputFile.Close()
	// 获取哈希值
	// Sum方法传递参数值nil，表示没有要合并的数据
	result := sha1.Sum(nil)
	// 打印结果
	fmt.Printf("文件%s的SHA1值为：\n%x\n", fileName, result)
}
