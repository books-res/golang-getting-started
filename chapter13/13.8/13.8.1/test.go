package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "stick"
	res := strings.HasPrefix(s, "st")
	fmt.Printf("“stick”是否以“st”开头？%t\n", res)

	s = "flush"
	res = strings.HasPrefix(s, "tr")
	fmt.Printf("“flush”是否以“tr”开头？%t\n", res)

	s = "photo"
	res = strings.HasSuffix(s, "ch")
	fmt.Printf("“photo”是否以“ch”结尾？%t\n", res)
}
