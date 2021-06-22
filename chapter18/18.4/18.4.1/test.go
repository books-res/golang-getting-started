package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	var myKey *rsa.PrivateKey
	myKey, _ = rsa.GenerateKey(rand.Reader, 1024)
	// 私钥已生成

	// 由于私钥中包含公钥的数据
	// 故可以通过PublicKey方法获取公钥
	var pubKey *rsa.PublicKey
	pubKey = &myKey.PublicKey

	fmt.Printf("私钥：%#v\n", myKey)
	fmt.Printf("\n公钥：%#v\n", pubKey)
}
