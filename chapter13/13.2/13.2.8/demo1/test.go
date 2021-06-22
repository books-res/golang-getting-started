package main

import "fmt"

func main() {
	var v = 13.50105485217

	fmt.Printf("精度为2：%.2f\n", v)
	fmt.Printf("精度为4：%.4f\n", v)
	fmt.Printf("精度为5：%.5f\n", v)
}
