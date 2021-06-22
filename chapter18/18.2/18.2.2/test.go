package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"math/rand"
)

func main() {
	dt := []byte("黄河之水天上来")
	// 确定数据的大小为BlockSize的整数倍
	var n int
	if len(dt) <= des.BlockSize {
		n = 1
	} else {
		n = len(dt) / des.BlockSize
		if len(dt)%des.BlockSize > 0 {
			n++
		}
	}
	srcdata := make([]byte, des.BlockSize*n)
	// 复制数据到新的[]byte实例中
	copy(srcdata, dt)

	// 密钥
	var key = []byte{
		0x02, 0xD3, 0x45, 0x87,
		0x32, 0x53, 0x7A, 0x95,
	}
	// 随机生成IV
	var iv = make([]byte, des.BlockSize)
	rand.Read(iv)

	// 创建Block对象
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 产生BlockMode对象
	encBlockmode := cipher.NewCBCEncrypter(block, iv)
	// 加密
	var out = make([]byte, len(srcdata))
	encBlockmode.CryptBlocks(out, srcdata)

	fmt.Printf("数据原文：%s\n", srcdata)
	fmt.Printf("加密后：%x\n", out)

	// 解密
	decBlockmode := cipher.NewCBCDecrypter(block, iv)
	var dec = make([]byte, len(out))
	decBlockmode.CryptBlocks(dec, out)
	fmt.Printf("解密后：%s\n", dec)
}
