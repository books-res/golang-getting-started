package main

import "fmt"

func main() {
	var i = 72
	switch {
	case i % 3 == 0:
		fmt.Printf("%d能被3整除\n", i)
	case i % 5 == 0:
		fmt.Printf("%d能被5整除\n", i)
	case i % 10 == 0:
		fmt.Printf("%d能被10整除\n", i)
	}
}