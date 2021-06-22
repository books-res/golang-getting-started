package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	// 准备密钥，本例中随机生成密钥
	// 生成私钥
	var prvKey, _ = rsa.GenerateKey(rand.Reader, 512)
	// 从私钥中获取公钥
	var pubKey = &prvKey.PublicKey

	// 待加密的消息
	var msg = "测试内容"
	fmt.Printf("原消息：%s\n", msg)

	// 加密
	var cipherText, _ = rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(msg))
	fmt.Printf("加密后：%x\n", cipherText)

	// 解密
	var decText, _ = rsa.DecryptPKCS1v15(rand.Reader, prvKey, cipherText)
	fmt.Printf("解密后：%s\n", decText)
}
