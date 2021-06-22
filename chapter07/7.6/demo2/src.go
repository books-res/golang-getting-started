package main

import "fmt"

func main() {
	func (who string) {
		fmt.Printf("Hello, %s\n", who)
	}("Jack")
}