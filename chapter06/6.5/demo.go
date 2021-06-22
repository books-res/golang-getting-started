package main

import "fmt"

const (
	A = iota
	B
	C
	D
)

const (
	H = iota + 1
	I
	J
	K
)


const (
	S = 17
	T = 9
	U = iota
	V
	W = iota + 3		// 7
	X					// 8
	Y					// 9
	Z					// 10
)

func main() {
	fmt.Printf("A: %d\nB: %d\nC: %d\nD: %d\n", A,B,C,D)
	fmt.Printf("\nH: %d\nI: %d\nJ: %d\nK: %d\n", H,I,J,K)
	fmt.Printf("\nS: %d\nT: %d\nU: %d\nV: %d\nW: %d\nX: %d\nY: %d\nZ: %d\n", S,T,U,V,W,X,Y,Z)
}