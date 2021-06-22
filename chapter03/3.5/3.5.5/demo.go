package main

import "fmt"

func main() {
	var r uint8
	var x uint8 = 0b_1011_1101
	var y uint8 = 0b_1110_0001
	r = x ^ y
	fmt.Printf("%08b ^ %08b: %08b\n", x, y, r)

	x = 0b_11110000
	y = 0b_00001111
	r = x ^ y
	fmt.Printf("%08b ^ %08b: %08b", x, y, r)
}