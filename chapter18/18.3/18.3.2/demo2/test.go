package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	// 数据原文
	var src = []byte("abcdefg--opqrst*+@")

	var (
		// SHA1
		sha1Sum = sha1.Sum(src)
		// SHA256
		sha256Sum = sha256.Sum256(src)
		// SHA384
		sha384Sum = sha512.Sum384(src)
		// SHA512
		sha512Sum = sha512.Sum512(src)
	)

	// 输出结果
	fmt.Printf("原文：%s\n", src)
	fmt.Printf("sha1校验码：%x\n", sha1Sum)
	fmt.Printf("sha256校验码：%x\n", sha256Sum)
	fmt.Printf("sha384校验码：%x\n", sha384Sum)
	fmt.Printf("sha512校验码：%x\n", sha512Sum)
}
