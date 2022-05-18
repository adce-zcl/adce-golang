package main

import "fmt"

// 全局变量
var isOk bool

func main() {
	var name string = "zhangsan"
	var age = 10

	var (
		a int
		b string
		c bool
	)
	// 只有函数内部的变量才可以用短变量声明
	d := "ok"
	fmt.Println(name, age, isOk)
	fmt.Println(a, b, c, d)
	const n = 100
	fmt.Println(n)

	// iota可理解为const语句块中的行索引
	const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	const (
		aa, bb = iota + 1, iota + 2 //1,2
		cc, dd                      //2,3
		ee, ff                      //3,4
	)

}
