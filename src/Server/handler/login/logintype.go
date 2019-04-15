package login

import "github.com/sowhyim/game/src/proto/login"

type Login struct {
	Login    string `xorm:"Login"`
	Password string `xorm:"LPassword"`
	Name     string `xorm:"Name"`
}

func (s *Login)TableName()string{
	return "t_Login"
}

func (s *Login)ToProto()*login.Request{
	return &login.Request{
		Login:s.Login,
		Password:s.Password,
	}
}