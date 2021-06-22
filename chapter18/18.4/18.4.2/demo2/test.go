package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 生成私钥
	var prvKey, _ = rsa.GenerateKey(rand.Reader, 1024)
	// 获取公钥
	var pubKey = &prvKey.PublicKey
	// 待加密的消息
	var msg = "pdr009@163.com"
	fmt.Printf("原消息：%s\n", msg)
	// 加密
	var cipherText, _ = rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, []byte(msg), nil)
	fmt.Printf("加密后：%x\n", cipherText)
	// 解密
	var decText, _ = rsa.DecryptOAEP(sha256.New(), rand.Reader, prvKey, cipherText, nil)
	fmt.Printf("解密后：%s\n", decText)
}
