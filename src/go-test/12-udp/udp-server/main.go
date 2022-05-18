// UDP server端
package main

import (
	"fmt"
	"net"
)

func main() {
	// 绑定IP地址
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 1227,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		// 接收数据，用字节数组接收，转换成string输出
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data:%v  addr:%v  count:%v\n", string(data[:n]), addr, n)
		// 用字节数组发送
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
