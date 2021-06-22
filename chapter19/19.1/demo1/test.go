package main

import (
	"fmt"
	"os"
)

func main() {
	var args = os.Args
	for index, val := range args {
		fmt.Printf("[%d]: %s\n", index, val)
	}
}
