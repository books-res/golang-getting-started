package test

type student struct {
	ID int
	Name string
	Age int
}

type Employee struct {
	ID int
	Name string
	Age int
	Phone int64
}

// 结构体自身可被外部代码访问，但其字段不可以
type Printer struct {
	current int16
	totalpage int64
	sn string
}

// 做以下修改后，字段成员可被外部代码访问
// type Printer struct {
// 	Current int16
// 	Totalpage int64
// 	Sn string
// }