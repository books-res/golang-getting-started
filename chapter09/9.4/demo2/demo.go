package main

import "fmt"

type iFile1 interface {
	getFilename() string
}

type iFile2 interface {
	iFile1					// 嵌套
	getTypeExt() string
}


type fileInfo struct {
	fn, ext string
}
// 实现接口的方法
func (f fileInfo) getFilename() string {
	return f.fn
}
func (f fileInfo) getTypeExt() string {
	return f.ext
}


func main() {
	var a iFile2 = fileInfo{"abcd.doc", ".doc"}
	fmt.Printf("文件名：%s\n扩展名：%s\n", a.getFilename(), a.getTypeExt())
	fmt.Println("-----------------------------")
	var b iFile1 = fileInfo{"city_list.txt", ".txt"}
	// 只能调用getFilename方法
	fmt.Printf("文件名：%s\n", b.getFilename())
}