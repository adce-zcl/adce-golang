package main

import (
	"fmt"
	"go-test/11-tcp/proto"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("dial failed,err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, how are you ? `
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed,err:", err)
			return
		}
		conn.Write(data)
	}
}
