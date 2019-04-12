package login

import (
	"Server/db"
	"context"
	"errors"
	"fmt"
	"proto/login"
)

type GameLoginServiceSrv struct{}

func NewServer() *GameLoginServiceSrv {
	return new(GameLoginServiceSrv)
}
func (s *GameLoginServiceSrv) GameLogin(ctx context.Context, req *login.LoginRequest) (*login.LoginReply, error) {
	var out Login
	fmt.Println(req.Login, req.Password)
	d := db.NewSession()
	_, err := d.Where("Login = ? and Password = ?", req.Login, req.Password).Get(&out)
	if err != nil {
		fmt.Println("login DB err :", err)
	}
	if out.Login == "" {
		return nil,errors.New("账号或密码错误，请确认！！")
	}
	fmt.Println(out)
	reply, err := db.GetRenWu(out.Name)
	if err != nil {
		fmt.Println("Get RenWu faild :", err)
	}
	return reply.ToProto(), err
}
