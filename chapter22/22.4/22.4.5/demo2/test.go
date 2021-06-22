package main

import (
	"fmt"
	"net/http"
)

var acs = map[string]string{
	"010":  "北京市",
	"022":  "天津市",
	"021":  "上海市",
	"0592": "福建省 厦门市",
	"027":  "湖北省 武汉市",
	"0794": "江西省 临川市",
	"0516": "江苏省 徐州市",
	"0571": "浙江省 杭州市",
	"0852": "贵州省 遵义市",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/qcode", func(w http.ResponseWriter, r *http.Request) {
		// 获取URL中的code参数
		var cv = r.URL.Query().Get("code")
		if cv == "" {
			w.Write([]byte("请提供要查询的区号"))
		} else {
			res, exist := acs[cv]
			if !exist {
				// 检索不到内容
				w.Write([]byte("您查询的区号暂时未收录"))
			} else {
				s := fmt.Sprintf("查询结果：%s", res)
				w.Write([]byte(s))
			}
		}
	})
	// 启动HTTP服务器
	http.ListenAndServe(":http", nil)
}
