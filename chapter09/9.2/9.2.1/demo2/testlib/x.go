package mylib

import "time"

type Order struct {
	ID uint64
	Product string
	Date time.Time
	Qty float32
	Remarks string
}

type Student struct {
	StdID uint
	Name string
	Age uint8
	email string		// 此字段只能在当前包中使用
}