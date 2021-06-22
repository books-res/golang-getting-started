package main

import "fmt"

func main() {
	var k uint8 = 0b_11011111
	var r = k &^ 0b_00000111
	fmt.Printf("%08b &^ 00000111: %08b\n", k, r)
}