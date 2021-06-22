package main

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	"fmt"
)

func main() {
	var msg = "A test message"
	// 获取各哈希算法相关的hash.Hash对象
	md5 := crypto.MD5.New()
	sha1 := crypto.SHA1.New()
	sha256 := crypto.SHA256.New()
	// 计算校验值
	var data = []byte(msg)
	md5.Write(data)
	sha1.Write(data)
	sha256.Write(data)
	fmt.Printf("原数据：%s\n", msg)
	fmt.Printf("MD5：%x\n", md5.Sum(nil))
	fmt.Printf("SHA1：%x\n", sha1.Sum(nil))
	fmt.Printf("SHA256：%x\n", sha256.Sum(nil))
}
