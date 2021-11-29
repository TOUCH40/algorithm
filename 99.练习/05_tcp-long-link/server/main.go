package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

type Manger struct {
	Kv map[int]*net.TCPConn
	fl int
}

func (m *Manger) run() {
	var dataNow int

	go func() {
		for {
			dataNow++
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		for k, v := range m.Kv {
			if dataNow%10 != k {
				fmt.Println("此时结果: ", dataNow%10)
				continue
			}

			_, err := v.Write([]byte(strconv.Itoa(dataNow)))
			if err != nil {
				log.Println(err)
			}

			time.Sleep(time.Millisecond * 300)
		}
	}
}

func main() {
	m := &Manger{Kv: make(map[int]*net.TCPConn)}

	go m.run()

	localAddress, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:3333") //定义一个本机IP和端口。
	tcpListener, err := net.ListenTCP("tcp", localAddress)        //在刚定义好的地址上进监听请求。
	if err != nil {
		fmt.Println("监听出错：", err)
		return
	}
	defer func() { //担心return之前忘记关闭连接，因此在defer中先约定好关它。
		tcpListener.Close()
	}()

	fmt.Println("ok")
	for {
		conn, err := tcpListener.AcceptTCP() //接受连接。
		if err != nil {
			fmt.Println("接受连接失败：")
			return
		}
		fmt.Println("收到链接")
		err = conn.SetKeepAlive(true)
		if err != nil {
			log.Println(err)
		}

		m.Kv[m.fl] = conn
		fmt.Println("收到订阅: ", m.fl)
		m.fl++
	}
}

func re(conn *net.TCPConn) {
	for {
		data := make([]byte, 1024*1024*10)
		read, err := conn.Read(data)
		if err != nil {
			//log.Println()
		}

		fmt.Println(string(data[:read]))
	}

}

func w(conn *net.TCPConn) {
	var i int
	i = 1
	for {
		_, err := conn.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			log.Println(err)
		}
		i++

		time.Sleep(time.Second * time.Duration(i))
	}
}
