package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 主页
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 返回index.htm文件的内容
		htmlFile, err := os.Open("index.htm")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		io.Copy(w, htmlFile)
		htmlFile.Close()
	})
	// 处理提交页
	http.HandleFunc("/setdata", func(w http.ResponseWriter, r *http.Request) {
		// 获取form字段的值
		var (
			name     = r.FormValue("name")
			phone    = r.FormValue("phone")
			bookname = r.FormValue("bookn")
			returned = r.FormValue("returned")
		)
		prtmsg := fmt.Sprintf("联系人：%s\n", name)
		prtmsg += fmt.Sprintf("手机号：%s\n", phone)
		prtmsg += fmt.Sprintf("书名：%s\n", bookname)
		if returned == "1" {
			prtmsg += "（此书已还）"
		} else {
			prtmsg += "（此书未还）"
		}
		prtmsg += "\n\n"
		fmt.Print("客户端提交的信息：\n")
		fmt.Println(prtmsg)
		w.Write([]byte("提交成功"))
	})
	// 启动HTTP服务
	http.ListenAndServe(":80", nil)
}
