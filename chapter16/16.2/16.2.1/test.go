package main

import (
	"errors"
	"fmt"
	"io"
)

type myBuffer struct {
	// 记录当前读写位置
	curPos int
	// 封装的内部数据缓冲区
	innerBuf []byte
}

// 实现读取功能
func (b *myBuffer) Read(p []byte) (n int, err error) {
	if p == nil || len(p) == 0 {
		n = 0
		err = nil
		return
	}
	if b.innerBuf == nil {
		err = errors.New("错误：未发现可用的数据")
		return
	}
	// 计算缓冲区剩余的字节数
	avalidBytes := len(b.innerBuf) - b.curPos
	if avalidBytes == 0 {
		// 可用字节数为0，表示已到达末尾
		err = io.EOF
		return
	}
	// 复制字节序列
	n = copy(p, b.innerBuf[b.curPos:])
	b.curPos += n
	return
}

// 实现写入功能
func (b *myBuffer) Write(p []byte) (n int, err error) {
	if p == nil || len(p) == 0 {
		n = 0
		err = nil
		return
	}
	// 如果内部缓冲区未初始化
	if b.innerBuf == nil {
		b.innerBuf = make([]byte, len(p))
		// 将当前位置移到开始位置
		b.curPos = 0
	}
	// 分析一下内部缓冲区是否可容纳新写入的数据
	avdSpc := cap(b.innerBuf) - len(b.innerBuf)
	if avdSpc < len(p) {
		// 重新分配空间
		newBuf := make([]byte, b.curPos+len(p))
		// 复制原缓冲区中的数据
		copy(newBuf, b.innerBuf[:b.curPos])
		// 替换旧的缓冲区
		b.innerBuf = newBuf
	}
	n = copy(b.innerBuf[b.curPos:], p)
	b.curPos += n
	return
}

// 复位函数，将当前位置移到数据流的开始处，以便调用Read方法
func (b *myBuffer) Reset() {
	b.curPos = 0
}

func CreateNewBuffer() *myBuffer {
	return &myBuffer{
		curPos:   0,
		innerBuf: make([]byte, 16),
	}
}

func main() {
	var bf = CreateNewBuffer()
	// 写入36个字节
	var data = make([]byte, 36)
	n := len(data)
	for i := 0; i < n; i++ {
		data[i] = byte(i + 1)
	}
	bf.Write(data)

	bf.Reset()
	fmt.Println("\n读出来的数据：")
	// 读取
	var temp = make([]byte, 5)
	for {
		c, err := bf.Read(temp)
		if err == io.EOF {
			// 已到末尾
			break
		}
		fmt.Println(temp[:c])
	}
}
