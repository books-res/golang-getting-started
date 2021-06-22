package main

import "fmt"

func main() {
	var n map[byte]string
	n[12] = "abc"
	n[24] = "opq"
	n[48] = "efg"
	fmt.Println(n)
}