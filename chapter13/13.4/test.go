package main

import "fmt"

type Product struct {
	Pid       uint
	ProdName  string
	ProdDate  string
	ProdColer string
}

func (p Product) String() string {
	return fmt.Sprintf("产品编号：%d，产品名称：%s，产品颜色：%s，生产日期：%s", p.Pid, p.ProdName, p.ProdColer, p.ProdDate)
}

func main() {
	var vp = Product{
		Pid:       41920014,
		ProdName:  "试验产品C",
		ProdColer: "白色",
		ProdDate:  "2017-9-27",
	}

	fmt.Println(vp)
	fmt.Printf("%v\n", vp)
	fmt.Printf("%s\n", vp)
}
