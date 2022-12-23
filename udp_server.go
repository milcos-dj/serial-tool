package main

import (
	"fmt"
	"net"
)

func main() {
	startServer()
}

func startServer() {
	lis, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 10001,
	})
	if err != nil {
		return
	}
	defer func() {
		err := lis.Close()
		if err != nil {
			fmt.Println("listen close error: ", err)
		}
	}()

	for {
		buf := make([]byte, 128)
		n, addr, err := lis.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ReadFromUDP error:", err)
			return
		}
		fmt.Printf("receive message %s from %v\n", string(buf[0:n]), addr)
	}
}
