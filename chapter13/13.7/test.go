package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "ABC#DEFG#HIJ#KLMN#OPQ#RST#UVW#XYZ"

	var res = strings.Split(s, "#")
	fmt.Printf("以“#”为隔符，拆分整个字符串：\n%s\n\n", strings.Join(res, "、"))

	res = strings.SplitN(s, "#", 2)
	fmt.Printf("以“#”为分隔符拆分字符串，但只拆分两次：\n%s\n\n", strings.Join(res, "、"))

	res = strings.SplitAfter(s, "#")
	fmt.Printf("在分隔符之后拆分，保留分隔符：\n%s\n\n", strings.Join(res, "、"))

	res = strings.SplitAfterN(s, "#", 3)
	fmt.Printf("在分隔符之后拆分，保留分隔符，仅拆分三次：\n%s\n", strings.Join(res, "、"))
}
