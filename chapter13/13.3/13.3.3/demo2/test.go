package main

import "fmt"

func main() {
	/*
		var y, m, d string
		fmt.Scanf("%s年%s月%s日", &y, &m, &d)

		fmt.Printf("y= %s, m= %s, d= %s\n", y, m, d)
	*/
	// 合理的做法
	var y, m, d int
	fmt.Scanf("%d年%d月%d日", &y, &m, &d)

	fmt.Printf("y= %d, m= %d, d= %d\n", y, m, d)
}
