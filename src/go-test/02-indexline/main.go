package main

import (
	"os"
)

func main() {
	buf := make([]byte, 1024)
	// 一般函数都是有多个返回值，第一个是真返回值，第二个是返回err
	// 如果不想处理err，就可以用_忽略这个返回值
	f, _ := os.Open("./file/text.txt")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
