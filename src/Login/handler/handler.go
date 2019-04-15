package handler

import (
	"context"
	pblogin "github.com/sowhyim/game/src/proto/login"
	"google.golang.org/grpc"
)

type GameLoginSrv struct{}

func (c *GameLoginSrv) CreateAccount(ctx context.Context, in *pblogin.CreateAccountRequest, opts ...grpc.CallOption) (*pblogin.CreateReply, error) {
	out := new(pblogin.CreateReply)

	return out, nil
}

func (c *GameLoginSrv) CreateName(ctx context.Context, in *pblogin.CreateRenWuRequest, opts ...grpc.CallOption) (*pblogin.CreateReply, error) {
	out := new(pblogin.CreateReply)

	return out, nil
}

func (c *GameLoginSrv) GameLogin(ctx context.Context, in *pblogin.LoginRequest, opts ...grpc.CallOption) (*pblogin.LoginReply, error) {
	out := new(pblogin.LoginReply)

	return out, nil
}
