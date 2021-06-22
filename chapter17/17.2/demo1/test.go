package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("fd01", 0644)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Mkdir("fd01/fd02", 0644)
	if err != nil {
		fmt.Println(err)
	}
}
