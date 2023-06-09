package main

import (
	"fmt"
	"net"
	"os"
)

var ipMap = make(map[string]int)

func main() {

	done := make(chan struct{})
	go startClient(done)

	<-done
}

func startClient(done chan struct{}) {

	laddr := &net.UDPAddr{
		IP:   net.ParseIP("192.168.8.104"),
		Port: 10002,
	}
	raddr := &net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 10001,
	}
	conn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		done <- struct{}{}
		return
	}
	if err != nil {
		done <- struct{}{}
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("", err)
		}
	}()
	go writePing(conn, raddr)
	readResp(conn)
}

func writePing(conn *net.UDPConn, raddr *net.UDPAddr) {
	for {
		var cmd string
		fmt.Print("请输入:")
		_, e := fmt.Scanf("%s", &cmd)
		if e != nil {
			fmt.Println("Scanf error:", e.Error())
			os.Exit(1)
		}
		if cmd == "s" {
			ipMap = make(map[string]int)
			_, err := conn.WriteToUDP([]byte(cmd), raddr)
			fmt.Println("WriteToUDP ping message")
			if err != nil {
				fmt.Printf("write ping message error: %v\n", err)
			}
		} else if cmd == "a" {
			fmt.Printf("get all device id:%v\n", ipMap)
		}

		_, _ = fmt.Scanln()
	}
}

func readResp(conn *net.UDPConn) {
	for {
		buf := make([]byte, 10)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ReadFromUDP error:", err)
			return
		}
		ipMap[addr.IP.String()] = addr.Port
		data := string(buf[0:n])
		fmt.Printf("receive message %s from %v\n", data, addr)
	}
}
