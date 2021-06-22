package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Date struct {
	year  int //年
	month int //月
	day   int //日
}

func (d Date) Year() int  { return d.year }
func (d Date) Month() int { return d.month }
func (d Date) Day() int   { return d.day }

func (d Date) String() string {
	return fmt.Sprintf("%d-%d-%d", d.year, d.month, d.day)
}

func (d *Date) Set(v string) error {
	parts := strings.Split(v, "/")
	if len(parts) != 3 {
		return errors.New("此值应该由三部分组成")
	}
	// 解析命令行参数值
	if y, err := strconv.ParseInt(parts[0], 10, 32); err != nil {
		return errors.New("Year 值不正确")
	} else {
		d.year = int(y)
	}
	if m, err := strconv.ParseInt(parts[1], 10, 32); err != nil {
		return errors.New("Month 值不正确")
	} else {
		d.month = int(m)
	}
	if dy, err := strconv.ParseInt(parts[2], 10, 32); err != nil {
		return errors.New("Day 值不正确")
	} else {
		d.day = int(dy)
	}
	return nil
}

func main() {
	// 绑定变量与命令行参数
	var thedate Date
	flag.Var(&thedate, "dt", "以“/”分隔的日期")
	// 分析命令行参数
	flag.Parse()
	// 打印结果
	fmt.Printf("日期：%d年%d月%d日\n", thedate.Year(), thedate.Month(), thedate.Day())
}
