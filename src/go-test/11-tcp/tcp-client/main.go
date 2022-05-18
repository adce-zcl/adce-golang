package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TCP client端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1226")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close()
	// 获取标准输入的对象
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 读取字符，知道遇到换行，并包含换行
		input, _ := inputReader.ReadString('\n')
		// 去掉换行回车
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		// 把字符串转换成字节流切片[]byte
		// 因为TCP是基于字节流的传输
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		// var buf [512]byte 等价
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed,err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
