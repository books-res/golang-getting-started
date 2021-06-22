package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nm uint32 = 10000
	theVal := reflect.ValueOf(nm)

	//fmt.Println(theVal)
	if theVal.Kind() == reflect.Uint32 {
		fmt.Printf("此对象的值：%v\n", theVal.Uint())
	}
}