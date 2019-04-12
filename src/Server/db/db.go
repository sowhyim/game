package db

import (
	"Server/enum"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	DB  *xorm.Engine
	URL string = "root:password@tcp(192.168.0.106:3306)/game?charset=utf8"
)

func InitDB() {
	var err error
	DB, err = xorm.NewEngine("mysql", URL)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}

func InsertRenWu(in enum.RenWu) error {
	c := NewSession()
	_, err := c.Insert(in)
	return err
}

func NewSession() *xorm.Session {
	session := DB.NewSession()
	return session
}
func GetAllRenWu() []*enum.RenWu {
	var out []*enum.RenWu
	db := NewSession()
	err := db.Where("IsOnline = true").Find(&out)
	if err != nil {
		panic(err)
	}
	db.Close()
	return out
}

func GetRenWu(Name string) (*enum.RenWu, error) {
	var out enum.RenWu
	db := NewSession()
	_, err := db.Where("Name = ?", Name).Get(&out)
	if err != nil {
		return nil, err
	}
	fmt.Println(out)
	db.Close()
	return &out, nil
}
