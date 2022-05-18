package main

import (
	"fmt"
	"math"
)

func main() {
	// 支持八进制、 六进制，以及科学记数法。标准库 math 定义了各数字类型取值范围。
	a, b, c, d := 071, 0x1F, 1e9, math.MinInt16
	fmt.Println(a, b, c, d)
	str := "你好"
	fmt.Println(str)

	// 转义字符
	fmt.Println("str := \"c:\\pprof\\main.exe\"")

	// 定义多行字符串
	// 反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	s1 := `第一行
    第二行
    第三行
    `
	s2 := "第一行\n第二行\n第三行"
	fmt.Println(s1)
	fmt.Println(s2)

	traversalString()
	changeString()
	sqrtDemo()
}

// 遍历字符串
func traversalString() {
	s := "pp.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	// 是复制数组，所以原字符串s1不会改变
	fmt.Println(string(byteS1))

	s2 := "博客"
	// 用于多字节字符的特殊类型
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

// go只有强制类型转换
func sqrtDemo() {
	a, b := 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
