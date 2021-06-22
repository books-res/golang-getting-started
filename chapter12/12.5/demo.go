package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

type User struct {
	Username string `maxlen:"10"`
	Password string `minlen:"8" mask:"*"`
}

func validateUser(u User) {
	var typofUser = reflect.TypeOf(u)
	var valofUser = reflect.ValueOf(u)
	// 读出Tag
	tagUn := typofUser.Field(0).Tag
	tagPwd := typofUser.Field(1).Tag
	// 解析用户名最大长度
	maxLenofUn, _ := strconv.ParseInt(tagUn.Get("maxlen"), 10, 0)
	// 解析密码最小长度
	minLenofPwd, _ := strconv.ParseInt(tagPwd.Get("minlen"), 10, 0)
	// 读取两个字段的值
	var usrName = valofUser.Field(0).String()
	var pssWd = valofUser.Field(1).String()
	// 开始验证
	unLen := utf8.RuneCountInString(usrName)
	pwdLen := utf8.RuneCountInString(pssWd)
	if unLen > int(maxLenofUn) {
		fmt.Printf("用户名 %s 长度超过%d\n", usrName, maxLenofUn)
		return
	}
	if pwdLen < int(minLenofPwd) {
		// 获取密码掩码
		maskChar := tagPwd.Get("mask")
		// 生成类似“********”的密码字符串
		displayPwd := strings.Repeat(maskChar, pwdLen)
		fmt.Printf("密码 %s 的长度不符，至少需要%d个字符\n", displayPwd, minLenofPwd)
		return
	}
	fmt.Println("用户信息符合要求")
}

func main() {
	var u1 User
	u1.Username = "jk__056546595xlase"
	u1.Password = "65656565656"
	validateUser(u1)

	fmt.Println()

	var u2 User
	u2.Username = "Jack"
	u2.Password = "321"
	validateUser(u2)

	fmt.Println()

	var u3 User
	u3.Username = "Alice"
	u3.Password = "UTU9ocg210359m82"
	validateUser(u3)
}
