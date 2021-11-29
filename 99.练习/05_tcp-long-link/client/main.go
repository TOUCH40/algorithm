package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	localAddress, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:3333") //定义一个本机IP和端口。
	dial, err := net.DialTCP("tcp", nil, localAddress)
	if err != nil {
		log.Println(err)
		return
	}
	dial.SetKeepAlive(true)

	for {
		data := make([]byte, 1024)
		read, err := dial.Read(data)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(string(data[:read]))
	}
}
