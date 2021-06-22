package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer = bytes.NewBuffer(nil)
	// 压缩三个文件
	gzw := gzip.NewWriter(buffer)
	for i := 1; i < 4; i++ {
		// 设置文件名
		gzw.Name = fmt.Sprintf("item-%d.txt", i)
		// 设置注释
		gzw.Comment = fmt.Sprintf("comment-#%d", i)
		// 写入数据
		gzw.Write([]byte(fmt.Sprintf("示例文本 %d\n", i)))
		// 关闭Writer对象
		gzw.Close()
		// 调用Restet方法很关键
		gzw.Reset(buffer)
	}

	// 解压缩
	gzr, err := gzip.NewReader(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 读取内容
	for {
		fmt.Printf("文件：%s\n", gzr.Name)
		fmt.Printf("注释：%s\n", gzr.Comment)
		fmt.Print("\n文件内容：\n")
		io.Copy(os.Stdout, gzr)
		fmt.Print("\n\n")
		// 下一个文件
		err := gzr.Reset(buffer)
		if err == io.EOF {
			// 到达流末尾，退出循环
			break
		}
	}
}
