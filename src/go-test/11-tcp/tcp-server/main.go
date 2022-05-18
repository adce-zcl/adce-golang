package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP server端

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		// 使用切片接收
		n, err := reader.Read(buf[:]) // 读数据
		if err != nil {
			fmt.Println("read from client failed,err:", err)
			break
		}
		// 读切片
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据:", recvStr)
		conn.Write([]byte(recvStr)) // 回射数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:1226")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		// 启动一个goroutine处理连接
		go process(conn)
	}
}
