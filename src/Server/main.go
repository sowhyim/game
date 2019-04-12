package main

import (
	"Server/db"
	"Server/env"
	"Server/handler/login"
	"google.golang.org/grpc"
	"net"
	pblogin "proto/login"
)

func main(){
	db.InitDB()
	env.ChangDiInit()
	lis,err:=net.Listen("tcp",":10001")
	defer lis.Close()
	if err!=nil{
		panic(err)
	}
	s:=grpc.NewServer()
	pblogin.RegisterGameLoginServiceServer(s,&login.GameLoginServiceSrv{})
	s.Serve(lis)
}