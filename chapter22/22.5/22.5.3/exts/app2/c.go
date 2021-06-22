package main

import "fmt"

func main() {
	fmt.Print("Content-Type: text/html;charset=utf-8\r\n")
	fmt.Print("\r\n")
	fmt.Print(`<html>
	<body>
		<p>欢迎来到【企业新闻】板块</p>
	</body>
	</html>`)
}
