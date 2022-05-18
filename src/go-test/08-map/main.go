package main

import "fmt"

func main() {
	// f2()
	f3()
}

func f1() {
	userinfo := map[string]string{
		"username": "bob",
		"password": "1234",
	}
	fmt.Println(userinfo)
	v, ok := userinfo["username"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("no")
	}

	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	delete(scoreMap, "小明") //将小明:100从map中删除
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}

func f2() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

func f3() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
