package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
)

var UdpClientHelp = flag.Bool("h", false, "帮助指令")
var InterfaceName = flag.String("in", "eno1", "网卡名称")

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
		time.Sleep(time.Second * 10)
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
func GetLocalIp() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	laddr := &net.UDPAddr{}
	for _, inter := range inters {
		fmt.Printf("addr name is : %v \n", inter.Name)
		if strings.Compare(inter.Name, *InterfaceName) == 0 {
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
