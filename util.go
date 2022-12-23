package main

import (
	"fmt"
	"net"
	"strings"
)

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
