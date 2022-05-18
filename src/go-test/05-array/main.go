package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 全局数组
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

// 多维数组
var arr00 [5][3]int
var arr11 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func main() {
	// f0()
	// f1()
	// f3()
	// f5()
	// f6()
	// f7()
	f8()
}

// 一维数组
func f0() {
	// 局部数组
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{2: 100, 4: 99}
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20}, // 最后一行要有逗号
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
}

// 多维数组
func f1() {
	a := [2][3]int{{1, 2, 3}, {4, 4, 4}}
	b := [...][2]int{{1, 2}, {222, 222}, {3, 3}}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(arr00)
	fmt.Println(arr11)
}

// 值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。
func f3() {
	a := [2]int{} // 全局数组可以不用{}
	fmt.Println(a)
	f3_test(a) // 有性能问题，值拷贝
	fmt.Println(a)
	f4(a)
}
func f3_test(x [2]int) {
	fmt.Printf("x:%p\n", &x)
	x[1] = 1000
}

// 返回数组长度
func f4(x [2]int) {
	fmt.Println(len(x), cap(x))
}

// 数组遍历
func f5() {
	var f [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for k1, v1 := range f { // k1 是 row， v1是一维数组
		for k2, v2 := range v1 { // k2 是col，v2是具体元素
			fmt.Printf("(%d,%d)=%d\t", k1, k2, v2)
		}
		fmt.Println()
	}
}

// 数组拷贝的传参
func f6() {
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [...]int{2, 3, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}
func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 求数组所有元素之和
func f7() {
	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	// rand.Seed(1)
	rand.Seed(time.Now().Unix()) // 设置随机数种子
	var b [10]int
	for i := 0; i < len(b); i++ {
		// 产生一个随机数
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	fmt.Printf("sum = %d\n", sum)
}
func sumArr(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

// 找出数组中和为给定值的两个下标
func f8() {
	b := [...]int{1, 2, 3, 4, 5}
	var target int = 4
	getSum(b, target)
}
func getSum(a [5]int, target int) {
	// 遍历数组
	for i := 0; i < len(a); i++ {
		target = target - a[i]
		for j := i + 1; j < len(a); j++ {
			if a[j] == target {
				fmt.Print("%d,%d\n", i, j)
			}
		}
	}
}
