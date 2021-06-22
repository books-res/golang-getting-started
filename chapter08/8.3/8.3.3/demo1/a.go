package main

import "fmt"

func main() {
	var n = 58
	switch {
	case n < 100:
		fmt.Println("该值小于100")
	case n < 80:
		fmt.Println("该值小于80")
	case n < 50:
		fmt.Println("该值小于50")
	case n < 30:
		fmt.Println("该值小于30")
	}
}