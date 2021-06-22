package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"math/rand"
	"os"
)

func main() {
	// 要加密的文件名
	var srcFilename = "测试文件.jpg"
	// 加密后的文件名
	var outFilename = "加密文件.bin"
	// 密钥 AES-192
	var key = []byte{
		0x18, 0x5F, 0x02, 0x20, 0x57, 0xB4, 0x13, 0x58,
		0x87, 0xDA, 0x4C, 0x6F, 0x39, 0x28, 0x77, 0x83,
		0x46, 0x9E, 0x74, 0x66, 0x74, 0x81, 0x52, 0x16,
	}
	// 初始向量，大小与BlockSize相同
	var iv = make([]byte, aes.BlockSize)
	rand.Read(iv)

	// -------- 加密 --------
	// 打开文件
	srcFile, err := os.Open(srcFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()
	// 创建加密后的文件
	outFile, err := os.Create(outFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	// 创建Block对象
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 将iv写入文件中，方便解密时读取
	outFile.Write(iv)
	// 创建Stream对象
	var stream = cipher.NewOFB(block, iv)
	// 创建StreamWriter实例
	strWriter := cipher.StreamWriter{
		S: stream,
		W: outFile,
	}
	defer strWriter.Close()
	// 把文件内容写入（加密）
	io.Copy(strWriter, srcFile)
}
