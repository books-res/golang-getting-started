package main

import "fmt"

func main() {
	var a int16 = 29156
	var r = a << 7
	fmt.Printf("%[1]d(%016[1]b) << 5: %[2]d(%016[2]b)", a, r)
}