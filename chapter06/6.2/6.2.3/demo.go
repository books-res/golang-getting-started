package main

import "fmt"

func main() {
	c1 := 15.545
	c2 := -0.0075
	fmt.Println(c1)
	fmt.Println(c2)

	// 科学记数法
	k1 := 1.4e6		//1400000.00
	fmt.Printf("%.2f\n", k1)

	k2 := 0.02E-3	//0.000020
	fmt.Printf("%f\n", k2)

	k3 := 100.55e+4		//1005500.0
	fmt.Printf("%.1f\n", k3)

	// 使用分隔符
	t1 := 36.2_5e6		//36250000.00
	t2 := 8_16e-3		//0.816
	t3 := 37_5.692_178_9e1_5	//375692178900000000
	fmt.Printf("%.2f\n", t1)
	fmt.Printf("%f\n", t2)
	fmt.Printf("%f\n", t3)

	// 下面代码会出错
	//var q1 float32 = 0.3e_3
	//var q2 = 12_e-3
	//fmt.Printf("%f\t%f\n", q1,q2)
}