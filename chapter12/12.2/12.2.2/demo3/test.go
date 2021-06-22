package main

import (
	"fmt"
	"reflect"
)

type Demo struct { }
func (d Demo) doWork() { }

func main() {
	var obj Demo
	t := reflect.TypeOf(obj)
	mt,ok := t.MethodByName("doWork")
	if ok {
		fmt.Printf("方法：%s\n", mt.Name)
	}
}