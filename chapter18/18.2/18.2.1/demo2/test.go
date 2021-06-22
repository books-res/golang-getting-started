package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	// 密钥
	var key = []byte{
		0x11, 0x2D, 0x64, 0x88, 0xA6, 0xFE, 0x18, 0x3C,
		0xE5, 0x51, 0x08, 0x9E, 0x63, 0x24, 0x72, 0x10,
		0xE3, 0x7C, 0x0B, 0x16, 0x73, 0x26, 0x17, 0x38,
		0x65, 0x6A, 0xD2, 0x48, 0x57, 0x9F, 0xC8, 0x13,
	}

	// 待加密的数据
	var srcData = make([]byte, aes.BlockSize)
	copy(srcData, []byte("你好，世界"))

	// 创建Block对象
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 加密
	fmt.Printf("原文：%s\n", srcData)
	var enc = make([]byte, len(srcData))
	block.Encrypt(enc, srcData)
	fmt.Printf("加密后：%x\n", enc)

	// 解密
	var dec = make([]byte, len(srcData))
	block.Decrypt(dec, enc)
	fmt.Printf("解密后：%s\n", dec)
}
