package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("DR01/DR02/hot")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.Remove("DR01/DR02")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.Remove("DR01")
	if err != nil {
		fmt.Println(err)
	}
}
