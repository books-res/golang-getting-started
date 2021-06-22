package main

import "fmt"

type email struct {
	from, to string
	subject  string
	senddate string
}

func main() {
	var msgs = []email{
		email{
			from:     "jack@126.com",
			to:       "dg02@126.net",
			subject:  "进度表",
			senddate: "2020-4-2",
		},
		email{
			from:     "abcd@test.org",
			to:       "let2915@21cn.com",
			subject:  "工序更改记录",
			senddate: "2020-3-17",
		},
		email{
			from:     "ok365@opengnt.cn",
			to:       "37025562@tov.edu.cn",
			subject:  "课程安排",
			senddate: "2019-12-24",
		},
		email{
			from:     "hack@123.net",
			to:       "6365056@qq.com",
			subject:  "最新地图",
			senddate: "2020-3-6",
		},
		email{
			from:     "chenjin@163.com",
			to:       "doudou@1211.org",
			subject:  "送货清单",
			senddate: "2019-7-12",
		},
	}

	fmt.Printf("%-16s%-23s%-20s%-16s\n", "发件人", "收件人", "标题", "发送日期")
	fmt.Println("---------------------------------------------------------------------------")
	for _, ml := range msgs {
		fmt.Printf("%-20s%-24s%-15s%-16s\n", ml.from, ml.to, ml.subject, ml.senddate)
	}
}
