package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":60001")
	if err != nil {
		os.Exit(1)
	}
	for {
		conn, err := lis.Accept()
		log.Printf("receive a connection!, %s\n", conn.RemoteAddr().String())
		if err != nil {
			log.Println("Accept error!", err)
		}
		go processConn(conn)

	}
}

func processConn(conn net.Conn) {
	log.Printf("start process conn!, %s\n", conn.RemoteAddr().String())
	r := bufio.NewReader(conn)
	connBytes := make([]byte, 22)
	_, err := io.ReadFull(r, connBytes)
	if err != nil {
		log.Println("read error !", err)
		conn.Close()
		return
	}
	log.Println("connBytes: ", connBytes)

	qai1 := []byte{0xFE, 0x04, 0x00, 0x00, 0x00, 0x01, 0x25, 0xC5}
	qai2 := []byte{0xFE, 0x04, 0x00, 0x01, 0x00, 0x01, 0x74, 0x05}
	qai3 := []byte{0xFE, 0x04, 0x00, 0x02, 0x00, 0x01, 0x84, 0x05}
	qai4 := []byte{0xFE, 0x04, 0x00, 0x03, 0x00, 0x01, 0xD5, 0xC5}
	//qaiall := []byte{0xFE, 0x04, 0x00, 0x00, 0x00, 0x06, 0x64, 0x07}

	//idCmd := []byte{0xFE, 0x04, 0x03, 0xEE, 0x00, 0x08, 0x85, 0xB2}
	//fmt.Println("write query id bytes", idCmd)
	//conn.Write(idCmd)

	//idResp := make([]byte, 21)
	//io.ReadFull(r, idResp)

	//fmt.Println("query id resp:", idResp)
	for {
		_, err := conn.Write(qai1)
		if err != nil {
			log.Println("write error !", err)
			break
		}
		readBytes := make([]byte, 7)
		_, err = io.ReadFull(r, readBytes)
		if err != nil {
			log.Println("read error !", err)
			break
		}
		log.Println("read ai1", readBytes)

		_, err = conn.Write(qai2)
		if err != nil {
			break
		}
		readBytes2 := make([]byte, 7)
		io.ReadFull(r, readBytes2)
		fmt.Println("read ai2", readBytes2)

		_, err = conn.Write(qai3)
		if err != nil {
			break
		}
		readBytes3 := make([]byte, 7)
		io.ReadFull(r, readBytes3)
		fmt.Println("read ai3", readBytes3)

		_, err = conn.Write(qai4)
		if err != nil {
			break
		}
		readBytes4 := make([]byte, 7)
		io.ReadFull(r, readBytes4)
		fmt.Println("read ai4", readBytes4)

		//conn.Write(qaiall)
		//readBytesAll := make([]byte, 7)
		//io.ReadFull(r, readBytesAll)
		//fmt.Println("read ai all", readBytesAll)

		time.Sleep(time.Second * 20)
	}

	log.Println("conn close!")
	conn.Close()

}
