package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

var Key []byte = make([]byte, 64)

func init() {
	// 生成密钥
	rand.Seed(time.Now().Unix())
	rand.Read(Key)
	fmt.Printf("生成的密钥：%x\n\n", Key)
}

func ComputeHMAC(msg string) []byte {
	var buf = []byte(msg)
	mHmac := hmac.New(sha256.New, Key)
	mHmac.Write(buf)
	return mHmac.Sum(nil)
}

func main() {
	// 第一条消息
	fmt.Print("请输入第一条消息：")
	var msg01 string
	fmt.Scanln(&msg01)
	// 计算HMAC
	var hash01 = ComputeHMAC(msg01)
	fmt.Printf("第一条消息的校验码：%x\n", hash01)
	// 第二条消息
	fmt.Print("请输入第二条消息：")
	var msg02 string
	fmt.Scanln(&msg02)
	// 计算HMAC
	var hash02 = ComputeHMAC(msg02)
	fmt.Printf("第二条消息的校验码：%x\n", hash02)
	// 验证两条消息是否一致
	res := hmac.Equal(hash01, hash02)
	if res {
		fmt.Println("\n两条消息一致")
	} else {
		fmt.Println("\n两条消息不一致")
	}
}
