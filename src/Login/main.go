package main

import (
	"fmt"
	"github.com/sowhyim/game/src/Login/listen"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":10002")
	if err != nil {
		panic(err)
	}
	i := 20
	for {
		if i <= 0 {
			continue
		}
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go listen.Dealer(conn, conn.RemoteAddr().String())
	}
}
