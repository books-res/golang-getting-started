package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type TextFileWriter struct {
	// 引用打开的文件
	file *os.File
}

func (w *TextFileWriter) WriteString(s string) (int, error) {
	if w.file == nil {
		return 0, errors.New("还没有打开文件")
	}
	return w.file.WriteString(s)
}
func (w *TextFileWriter) Close() error {
	return w.file.Close()
}

// 用来创建TextFileWriter实例
func CreateWriter(filename string) *TextFileWriter {
	file, err := os.Create(filename)
	if err != nil {
		return nil
	}
	return &TextFileWriter{file: file}
}

//------------------------------------------
type TextFileReader struct {
	file *os.File
}

func (r *TextFileReader) ReadString() (string, error) {
	if r.file == nil {
		return "", errors.New("还没有打开文件")
	}
	// 组建字符串
	var bd = new(strings.Builder)
	var bf = make([]byte, 32)
	for {
		n, err := r.file.Read(bf)
		if err == io.EOF {
			break
		}
		bd.Write(bf[:n])
	}
	// 返回字符串
	return bd.String(), nil
}
func (r *TextFileReader) Close() error {
	return r.file.Close()
}

// 用来创建TextFileReader实例
func CreateReader(filename string) *TextFileReader {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	return &TextFileReader{file: file}
}

//-------------------------------------

func main() {
	const fileName = "notes.txt"

	// 写入文本
	var wt = CreateWriter(fileName)
	wt.WriteString("桃花帘外东风软\n")
	wt.WriteString("桃花帘内晨妆懒\n")
	wt.WriteString("帘外桃花帘内人\n")
	wt.WriteString("人与桃花隔不远\n")
	// 关闭文件
	wt.Close()

	//---------------------------------
	// 读取文本
	var rd = CreateReader(fileName)
	var content, err = rd.ReadString()
	if err != nil {
		fmt.Println("错误：", err)
		rd.Close()
		return
	}
	// 关闭文件
	rd.Close()
	fmt.Printf("从文件中读到的文本：\n%s", content)
}
