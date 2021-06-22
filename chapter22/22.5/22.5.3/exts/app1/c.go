package main

import "fmt"

func main() {
	// 写入HTTP消息头
	fmt.Print("Content-Type: text/html;charset=utf-8\r\n")
	fmt.Print("\r\n") // 分隔行
	// 写入HTTP消息正文
	fmt.Print(
		`<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8" />
		<title>Test Page</title>
	</head>
	<body>
		<p>欢迎来到【产品信息】板块</p>
	</body>
</html>
`)
}
