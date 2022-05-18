// 客户端
package main

import (
	"fmt"
	"net"
)

func main() {
	// 连接
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 1227,
	})
	if err != nil {
		fmt.Println("连接服务器失败,err:", err)
		return
	}
	defer socket.Close()
	// 初始化一个切片
	sendData := []byte("Hello server")
	// 发送切片
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("发送数据失败,err:", err)
		return
	}
	// 构建一个最大容量为4096的切片
	data := make([]byte, 4096)
	// 将数据接收在切片里
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("接收数据失败,err:", err)
		return
	}
	fmt.Printf("recv:%v  addr:%v  count:%v\n", string(data[:n]), remoteAddr, n)
}
