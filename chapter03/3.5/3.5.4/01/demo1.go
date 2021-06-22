package main

import "fmt"

func main() {
	var a uint8 = 0b_1001_1101
	var r = a >> 3
	fmt.Printf("%08b >> 3: %08b\n", a, r)
	a = 0b_1010_1001
	r = a << 2
	fmt.Printf("%08b << 2: %08b", a, r)
}