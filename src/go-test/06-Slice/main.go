// 1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
// 2. 切片的长度可以改变，因此，切片是一个可变的数组。
// 3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
// 4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
// 5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
// 6. 如果 slice == nil，那么 len、cap 结果都等于 0。
package main

import "fmt"

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
func main() {
	// f1()
	// f2()
	f3()
}

// 切片的创建
func f1() {
	// 1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不为空")
	}
	fmt.Println("s1:", s1)

	// 2. :=
	s2 := []int{}
	fmt.Println("s2:", s2)

	// 3.make()
	var s3 []int = make([]int, 3)
	fmt.Println("s3:", s3)

	// 4.初始化赋值
	var s4 []int = make([]int, 1, 100)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	// 5.从数组中切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	s6 = arr[1:4] // 前包后不包[1,4)
	fmt.Println(s6)
}

func f2() {
	fmt.Printf("全局变量：arr %v\n", arr)
	fmt.Printf("全局变量：slice0 %v\n", slice0)
	fmt.Printf("全局变量：slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)
	fmt.Printf("-----------------------------------\n")
	arr2 := [...]int{90, 80, 70, 60, 50, 40, 30, 20, 10, 0}
	slice5 := arr2[2:8]
	slice6 := arr2[0:6]         //可以简写为 slice := arr[:end]
	slice7 := arr2[5:10]        //可以简写为 slice := arr[start:]
	slice8 := arr2[0:len(arr)]  //slice := arr[:]
	slice9 := arr2[:len(arr)-1] //去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)
}

func f3() {
	data := [...]int{1, 2, 3, 4, 5, 10: 10}
	fmt.Println(data)
	s := data[:4:10]
	fmt.Println(s)
}
