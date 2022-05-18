package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

// Class 班级
type Class struct {
	Title   string
	Student []*Student
}

func main() {
	// f1()
	// f3()
	// f4()
	f5()
}

func f1() {
	s1 := Student{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"女"}
}

func f2() {
	c := Class{
		Title:   "101",
		Student: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     1,
		}
		c.Student = append(c.Student, stu)
	}

	// JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json:%s\n\n\n", data)

	// JSON反序列化
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	ce[1].age = 999
}
func f3() {
	var ce []student
	ce = []student{
		student{1, "xiaoming", 22},
		student{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
}

func f4() {
	ce := make(map[int]student)
	ce[1] = student{1, "xiaolizi", 22}
	ce[2] = student{2, "wang", 23}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
}

// 实现map的有序输出
func f5() {
	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"

	sli := []int{}
	for k, _ := range map1 {
		sli = append(sli, k)
	}
	sort.Ints(sli)
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[sli[i]])
	}

}
