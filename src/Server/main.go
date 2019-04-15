package main

import (
	"github.com/sowhyim/game/src/Server/env"
	"github.com/sowhyim/game/src/Server/handler/login"
	"github.com/sowhyim/game/src/gamedb"
	pblogin "github.com/sowhyim/game/src/proto/login"
	"google.golang.org/grpc"
	"net"
)

func main() {
	db.InitDB()
	env.ChangDiInit()
	lis, err := net.Listen("tcp", ":10001")
	defer lis.Close()
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pblogin.RegisterGameLoginServiceServer(s, &login.GameLoginServiceSrv{})
	s.Serve(lis)
}
