package listen

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sowhyim/game/src/enum"
	"github.com/sowhyim/game/src/proto/pbLogin"
	"google.golang.org/grpc"
	"net"
)

type ClientConn map[string]net.Conn

var Conn ClientConn = ClientConn{}

func Dealer(c net.Conn, index string) {
	Conn[index] = c
	user := GameLogin(index)
	if user != nil {
		return
	}
	index = user.Name
	Conn[index] = c
	fmt.Println(Conn)
	HandleScan(index, user)
	Conn[index].Close()
	delete(Conn, index)
}

func GameLogin(index string) *enum.RenWu {
	var login, password string
LOOP:
	var by = make([]byte, 1024)
	Conn[index].Write([]byte(enum.LoginText))
	Conn[index].Read(by)
	login = string(by)
	fmt.Println(string(by))
	//var by2 = make([]byte, 1024)
	Conn[index].Write([]byte(enum.PasswordTest))
	Conn[index].Read(by)
	password = string(by)
	fmt.Println(string(by))
	user, err := CallServer(login, password)
	if err != nil {
		Conn[index].Write([]byte(err.Error()))
		goto LOOP
	}
	Conn[index].Write([]byte(enum.LoginSuccess))
	usermsg, err := json.Marshal(user)
	Conn[index].Write([]byte(usermsg))
	if Conn[user.Name] != nil {
		Conn[index].Write([]byte(enum.UserLoginError))
		Conn[index].Close()
		return nil
	}
	delete(Conn, index)
	return user
}

func HandleScan(index string, user *enum.RenWu) {
	for {
		Conn[index].Write([]byte("你可以输入以下操作：1、移动，2、使用道具，3、退出"))
		by := make([]byte, 1024)
		Conn[index].Read(by)
		if string(by) == "1" || string(by) == "移动" {
			fmt.Println("by = 1")
		}
		if string(by) == "2" || string(by) == "使用道具" {
			fmt.Println("by = 2")
		}
		switch string(by) {
		case "1":
			Conn[index].Write([]byte("请移动"))
			fmt.Println("移动")
		case "移动":
			Conn[index].Write([]byte("请移动"))
			fmt.Println("移动")
		case "2":
			Conn[index].Write([]byte("请选择需要使用的道具"))
			fmt.Println("请选择需要使用的道具")
		case "使用道具":
			Conn[index].Write([]byte("请选择需要使用的道具"))
			fmt.Println("请选择需要使用的道具")
		case "3":
			return
		case "退出":
			return
		}
	}
}

func CallServer(login, password string) (*enum.RenWu, error) {
	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	client := pbLogin.NewGameLoginServiceClient(conn)
	out, err := client.GameLogin(context.Background(), &pbLogin.LoginRequest{Login: login, Password: password})
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, errors.New("请确认你的账号或者密码")
	}

	return enum.FromProto(out), nil
}
