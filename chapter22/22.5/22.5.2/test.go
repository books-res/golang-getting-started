package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var cmd = &exec.Cmd{
			Path: "test",
			Dir:  "./cgi-bin",
		}
		// 获取标准输出通道
		var outputPipe, _ = cmd.StdoutPipe()
		err := cmd.Start()
		defer cmd.Wait()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 读取输出内容
		outReader := bufio.NewReaderSize(outputPipe, 256)
		// 读HTTP消息头
		for {
			line, prefix, err := outReader.ReadLine()
			if prefix {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err == io.EOF {
				break
			}
			// 遇到空白行，表示消息头部分已结束
			if len(line) == 0 {
				break
			}
			var header = strings.SplitN(string(line), ":", 2)
			if len(header) < 2 {
				// 该行消息头不完整，跳过
				continue
			}
			headername := strings.TrimSpace(header[0])
			headerval := strings.TrimSpace(header[1])
			w.Header().Add(headername, headerval)
		}
		// 读取消息正文
		_, err = io.Copy(w, outReader)
		if err != nil {
			fmt.Println(err)
			cmd.Process.Kill()
		}
	})

	http.ListenAndServe(":http", nil)
}

/*
服务器响应的消息头：
Access-Token: 52fdfc072182654f163f5f0f9a621d
Content-Length: 41
Content-Type: text/html; charset=utf-8
Server-Ver: 1.0

响应消息正文：
<h3>这是一个简单的 CGI 程序</h3>
*/
