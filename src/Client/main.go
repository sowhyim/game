package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

type ClientConn map[int]net.Conn

var Conn ClientConn = ClientConn{}
var WChan = make(chan int)

var URL = "localhost:10002"

func main() {
	conn, err := net.Dial("tcp", URL)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	Conn[1] = conn
	go Output(Conn)
	go Input(Conn)
	WChan <- 1
}

func Output(c ClientConn) {
	for {
		var by = make([]byte, 1024)
		var byZero = make([]byte, 0)
		n,err:=c[1].Read(by)
		if err!=nil{
			fmt.Println(err)
		}
		data:= by[0:n]
		fmt.Println(len(data))
		switch string(data) {
		case "该用户已经登陆！":
			fmt.Println("该用户已经登陆")
			c[1].Close()
			delete(c, 1)
			<-WChan
		case string(byZero):
			ReConn()
		default:
			fmt.Println(string(data))
		}

	}
}

func Input(c ClientConn) {
	for {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		c[1].Write([]byte(input.Text()))
	}
}

func ReConn() {
	Conn[1].Close()
	delete(Conn, 1)
	for i := 3; i > 0; i-- {
		fmt.Println("连接断开，重新连接中。。。")
		conn, err := net.Dial("tcp", URL)
		if err == nil {
			Conn[1] = conn
			return
		}
		time.Sleep(time.Second)
	}
	if Conn[1] != nil {
		fmt.Println("连接成功！")
	} else {
		fmt.Println("连接失败！")
		<-WChan
	}
}
