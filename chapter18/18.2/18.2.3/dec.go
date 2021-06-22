package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
)

func main() {
	// 已加密的文件名
	var encFilename = "加密文件.bin"
	// 解密后的文件名
	var decFilename = "解密后.jpg"
	// 密钥，必须与加密时使用的一致
	var key = []byte{
		0x18, 0x5F, 0x02, 0x20, 0x57, 0xB4, 0x13, 0x58,
		0x87, 0xDA, 0x4C, 0x6F, 0x39, 0x28, 0x77, 0x83,
		0x46, 0x9E, 0x74, 0x66, 0x74, 0x81, 0x52, 0x16,
	}

	// -------- 解密 --------
	// 打开已加密的文件
	encFile, err := os.Open(encFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建解密后的文件
	decFile, err := os.Create(decFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer encFile.Close()
	defer decFile.Close()
	// 创建Block对象
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 从文件中读出iv
	var iv = make([]byte, aes.BlockSize)
	encFile.Read(iv)
	// 创建Stream对象
	stream := cipher.NewOFB(block, iv)
	// 创建StreamReader实例
	strReader := cipher.StreamReader{
		S: stream,
		R: encFile,
	}
	io.Copy(decFile, strReader)
}
