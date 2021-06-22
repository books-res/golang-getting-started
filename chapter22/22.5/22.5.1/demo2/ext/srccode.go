package main

import (
	"fmt"
	"os"
)

func main() {
	// 读取环境变量
	appName, exist := os.LookupEnv("APP_NAME")
	if !exist {
		fmt.Print("未找到 APP_NAME 环境变量")
		return
	}
	appID, exist := os.LookupEnv("APP_ID")
	if !exist {
		fmt.Print("未找到 APP_ID 环境变量")
		return
	}
	devID, exist := os.LookupEnv("DEV_ID")
	if !exist {
		fmt.Print("未找到 DEV_ID 环境变量")
		return
	}
	// 输出
	fmt.Print("Content-Type: text/html; charset=utf-8\r\n")
	fmt.Print("\r\n")
	fmt.Printf(
		`<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>测试页面</title>
  </head>
  <body>
	<h3>应用名称：</h3>
	<p>%s</p>
	<h3>应用标识：</h3>
	<p>%s</p>
	<h3>开发者标识：</h3>
	<p>%s</p>
  </body>
</html>`, appName, appID, devID)
}
