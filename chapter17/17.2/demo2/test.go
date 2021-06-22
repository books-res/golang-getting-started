package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("fd01/fd02", 0644)
	if err != nil {
		fmt.Println(err)
	}
}
