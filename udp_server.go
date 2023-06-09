package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	startServer()
}

func startServer() {
	conn, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 10001,
	})
	if err != nil {
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("listen close error: ", err)
		}
	}()
	for {
		buf := make([]byte, 4096)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ReadFromUDP error:", err)
			return
		}
		data := string(buf[0:n])
		fmt.Printf("receive message %s from %v\n", data, addr)
		if data == "s" {
			_, _ = conn.WriteToUDP([]byte("ok"), addr)
		}

	}
}

func GetLocalIp(InterfaceName string) string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	laddr := &net.UDPAddr{}
	for _, inter := range inters {
		fmt.Printf("addr name is : %v \n", inter.Name)
		if strings.Compare(inter.Name, InterfaceName) == 0 {
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//判断是否存在IPV4 IP 如果没有过滤
					ip := ipnet.IP.To4()
					if ip != nil {
						fmt.Printf("addr name s%, ip is %s \n", inter.Name, ip.String())
						laddr.IP = net.ParseIP(ip.String())
						return ip.String()
					}
				}
			}
		}
	}
	return ""
}
