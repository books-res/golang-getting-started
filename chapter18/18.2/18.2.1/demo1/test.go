package main

import (
	"crypto/des"
	"fmt"
)

func main() {
	// 密钥
	var key = []byte{
		0x2D, 0x11, 0x25, 0xA4,
		0x8E, 0x74, 0x60, 0x13,
	}
	// 创建对应的Block对象
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 数据明文
	var test = []byte{
		0x26, 0x23, 0xF7, 0xA2,
		0x29, 0x33, 0xC4, 0x45,
		0x72, 0x19, 0x9B, 0x42,
	}
	// 用于存放密文，大小与明文相同
	var enc = make([]byte, len(test))
	// 加密
	block.Encrypt(enc, test)
	fmt.Printf("加密前：%x\n", test)
	fmt.Printf("加密后：%x\n", enc)
	// 解密
	var dec = make([]byte, len(test))
	block.Decrypt(dec, enc)
	fmt.Printf("解密后：%x\n", dec)
}
