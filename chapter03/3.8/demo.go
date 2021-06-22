package main

import "fmt"

func f1() {
	var z = 0.01
	z = z + 0.02
	fmt.Printf("z = z + 0.02: %f\n", z)
}

func f2() {
	var z = 0.01
	z += 0.02
	fmt.Printf("z += 0.02: %f\n", z)
}

func main() {
	f1()
	f2()
}