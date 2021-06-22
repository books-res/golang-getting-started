package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

func main() {
	// 生成密钥
	var prvkey, _ = rsa.GenerateKey(rand.Reader, 512)
	// 从私钥中取出公钥
	var pubkey = &prvkey.PublicKey
	// 导出二进制数据
	// 仅转化私钥
	keydata := x509.MarshalPKCS1PrivateKey(prvkey)
	// 把它写入到文件中
	ioutil.WriteFile("myPrvkey", keydata, 0600)

	fmt.Println("密钥已存入文件")
	// 为了验证密钥是否已正确存储
	// 先用公钥将消息加密，稍后从文件中读取私钥来解密
	var testmsg = "测试消息"
	var cipherdata, _ = rsa.EncryptPKCS1v15(rand.Reader, pubkey, []byte(testmsg))
	fmt.Printf("%s 被加密后：\n%x\n\n", testmsg, cipherdata)

	// 从文件中读出密钥（私钥）
	indata, err := ioutil.ReadFile("myPrvkey")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 加载密钥
	loadKey, err := x509.ParsePKCS1PrivateKey(indata)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("已从文件中加载密钥")

	// 尝试解密
	decrdata, err := rsa.DecryptPKCS1v15(rand.Reader, loadKey, cipherdata)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("被解密的消息：%s\n", decrdata)
}
