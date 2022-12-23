package main

import (
	"fmt"
	"go.bug.st/serial"
	"os"
)

func main() {
	ports, err := serial.GetPortsList()
	if err != nil {
		os.Exit(1)
	}
	if len(ports) == 0 {
		os.Exit(1)
	}

	for _, port := range ports {
		fmt.Printf("串口号：%s\n", port)
	}
}
