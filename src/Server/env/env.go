package env

import (
	"fmt"
	"Server/db"
	"Server/enum"
)

type ChangDi struct {
	Map *Map
	List []*enum.RenWu
}

var InitChangDi ChangDi

func ChangDiInit(){
	InitChangDi.Map=&Map{100,100}
	InitChangDi.List = db.GetAllRenWu()
	for _,val := range InitChangDi.List{
		fmt.Println(val)
	}
}

func CreateRenWu(name string, sex int32, race int32) *enum.RenWu {
	var initRenWu = enum.InitRace(race)
	initRenWu.Name = name
	initRenWu.Sex = sex
	return initRenWu
}

type Map struct {
	X int32
	Y int32
}
