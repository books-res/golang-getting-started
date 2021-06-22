package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.RemoveAll("DR01")
	if err != nil {
		fmt.Println(err)
	}
}
