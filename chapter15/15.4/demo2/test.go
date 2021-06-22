package main

import (
	"fmt"
	"sort"
	"time"
)

type goods struct {
	id   uint32 // 商品编号
	name string // 商品名称
	qty  uint16 // 商品数量
	time int64  // 入库时间
}

type goodsCollection []goods

func (c goodsCollection) Len() int {
	return len(c)
}
func (c goodsCollection) Less(i, j int) bool {
	return c[i].time < c[j].time
}
func (c goodsCollection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	var mygoods = goodsCollection{
		goods{
			id:   7201,
			name: "商品A",
			qty:  18,
			time: time.Date(2019, time.May, 11, 12, 4, 12, 0, time.Local).Unix(),
		},
		goods{
			id:   7202,
			name: "商品B",
			qty:  200,
			time: time.Date(2018, time.September, 21, 16, 30, 22, 0, time.Local).Unix(),
		},
		goods{
			id:   7203,
			name: "商品C",
			qty:  102,
			time: time.Date(2020, time.January, 2, 17, 3, 10, 0, time.Local).Unix(),
		},
		goods{
			id:   7204,
			name: "商品D",
			qty:  70,
			time: time.Date(2019, time.June, 20, 15, 45, 0, 0, time.Local).Unix(),
		},
		goods{
			id:   7205,
			name: "商品E",
			qty:  65,
			time: time.Date(2017, time.November, 28, 9, 16, 13, 0, time.Local).Unix(),
		},
		goods{
			id:   7206,
			name: "商品F",
			qty:  178,
			time: time.Date(2019, time.March, 11, 14, 37, 10, 0, time.Local).Unix(),
		},
	}

	// 排序前输出
	fmt.Println("--------- 排序前 ---------")
	printGoods(mygoods)

	// 排序
	sort.Sort(mygoods)

	// 排序后输出
	fmt.Println("\n--------- 排序后 ---------")
	printGoods(mygoods)
}

func printGoods(d goodsCollection) {
	for _, g := range d {
		fmt.Printf("%d - %s\t入库时间戳：%d\n", g.id, g.name, g.time)
	}
}
