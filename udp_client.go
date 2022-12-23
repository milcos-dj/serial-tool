package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var UdpClientHelp = flag.Bool("h", false, "帮助指令")
var InterfaceName = flag.String("in", "eth0", "网卡名称")

func main() {
	flag.Parse()
	if *UdpClientHelp {
		flag.Usage()
		return
	}
	for {
		err := startClient()
		if err != nil {
			continue
		}
		time.Sleep(time.Minute)
	}
}

func startClient() error {
	laddr := &net.UDPAddr{
		IP: net.ParseIP(GetLocalIp()),
	}
	raddr := &net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 10001,
	}
	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		return err
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("")
		}
	}()
	_, err = conn.Write([]byte("ping"))
	fmt.Println("write ping message")
	if err != nil {
		return err
	}
	return err
}
