package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tarm/serial"
	"io"
	"log"
)

type Config struct {
	Name string
	Baud int
	DataBits byte
	Parity int
	StopBits int
	DataSize int
}

//export StartPort
func StartPort(cnf *Config) {
	conf := &serial.Config{
		Name: cnf.Name,
		Baud: cnf.Baud,
		Size: cnf.DataBits,
		Parity: getParity(cnf.Parity),
		StopBits: getStopBit(cnf.StopBits),
	}

	port, err := serial.OpenPort(conf)
	if err != nil {
		log.Panic("serial open port error!")
	}
	readData(port, cnf)
}


func getParity(parity int) serial.Parity{
	var p = serial.ParityNone
	if parity == 1 {
		p = serial.ParityOdd
	} else if parity == 2 {
		p = serial.ParityEven
	} else if parity == 3 {
		p = serial.ParityMark
	} else if parity == 4 {
		p = serial.ParitySpace
	}
	return p
}

func getStopBit(stopBit int) serial.StopBits{
	var p = serial.Stop1
	if stopBit == 1 {
		p = serial.Stop1Half
	} else if stopBit == 2 {
		p = serial.Stop2
	}
	return p
}

func readData(port *serial.Port, cnf *Config) {
	r := bufio.NewReaderSize(port, 1024)
	defer port.Close()
	for {
		var buf []byte
		buf = make([]byte, cnf.DataSize)
		//port.Read(buf)
		io.ReadFull(r, buf)
		log.Println(buf)
	}

}

var Help = flag.Bool("h", false, "帮助指令")
var Name = flag.String("n", "COM3", "串口号")
var Baud = flag.Int("b", 9600, "波特率")
var Stop = flag.Int("s", 0, "停止位")
var ParityNone = flag.Int("p", 0, "校验位")
var DataBits = flag.Int("d", 8, "数据位")
var DataSize = flag.Int("t", 12, "每帧数据包大小")


func main(){

	flag.Parse()
	if *Help {
		flag.Usage()
		return
	}
	var cnf = &Config{
		Name:     *Name,
		Baud:     *Baud,
		StopBits: *Stop,
		Parity:   *ParityNone,
		DataBits: byte(*DataBits),
		DataSize:     *DataSize,
	}
	fmt.Println(*Name)
	fmt.Println(*Baud)
	fmt.Println(cnf)
	StartPort(cnf)

}