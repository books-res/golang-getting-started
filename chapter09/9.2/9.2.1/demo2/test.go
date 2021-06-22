package main

import (
	"fmt"
	"time"
	"./testlib"
)

func main() {
	var x mylib.Order = mylib.Order {
		ID: 10021,
		Product: "产品 A",
		Date: time.Now(),
		Qty: 5000.56,
		Remarks: "nothing",
	}
	fmt.Printf("%+v", x)
}