package enum

import pblogin "github.com/sowhyim/game/src/proto/login"

const (
	Human  = 1
	Elf    = 2
	Demon  = 3
	Dragon = 4
	Orc    = 5
)

func InitRace(sel int32) *RenWu {
	switch sel {
	case Human:
		return &HumanInit
	case Elf:
		return &ElfInit
	case Demon:
		return &DemonInit
	case Dragon:
		return &DragenInit
	case Orc:
		return &OrcInit
	}
	return nil
}

var HumanInit = RenWu{
	Life:            100,
	Magic:           100,
	Physical:        2,
	Sorcery:         4,
	DefensePhysical: 3,
	DefenseSorcery:  3,
}

var ElfInit = RenWu{
	Life:            80,
	Magic:           120,
	Physical:        1,
	Sorcery:         5,
	DefensePhysical: 1,
	DefenseSorcery:  5,
}

var DemonInit = RenWu{
	Life:            120,
	Magic:           80,
	Physical:        4,
	Sorcery:         2,
	DefensePhysical: 4,
	DefenseSorcery:  3,
}

var DragenInit = RenWu{
	Life:            200,
	Magic:           30,
	Physical:        5,
	Sorcery:         3,
	DefensePhysical: 2,
	DefenseSorcery:  4,
}

var OrcInit = RenWu{
	Life:            150,
	Magic:           50,
	Physical:        2,
	Sorcery:         4,
	DefensePhysical: 5,
	DefenseSorcery:  1,
}

type RenWu struct {
	Name            string `xorm:"RName"`
	Sex             int32  `xorm:"Sex"`
	Life            int32  `xorm:"Life"`
	Magic           int32  `xorm:"Magic"`
	Physical        int32  `xorm:"Physical"`
	Sorcery         int32  `xorm:"Sorcery"`
	DefensePhysical int32  `xorm:"DefensePhysical"`
	DefenseSorcery  int32  `xorm:"DefenseSorcery"`
	IsOnline        bool   `xorm:"IsOnline"`
	*ZhuangTai      `xorm:"-"`
	X               int32 `xorm:"X"`
	Y               int32 `xorm:"Y"`
}

type ZhuangTai struct {
}

type WeiZhi struct {
	X int32
	Y int32
}

func (a *RenWu) WeiZhi() *WeiZhi {
	return &WeiZhi{
		a.X,
		a.Y,
	}
}

func (c *RenWu) TableName() string {
	return "t_RenWu"
}

func (c *RenWu) ToProto() *pblogin.LoginReply {
	return &pblogin.LoginReply{
		Name:            c.Name,
		Sex:             c.Sex,
		Life:            c.Life,
		Magic:           c.Magic,
		Physical:        c.Physical,
		Sorcery:         c.Sorcery,
		DefensePhysical: c.DefensePhysical,
		DefenseSorcery:  c.DefenseSorcery,
		X:               c.X,
		Y:               c.Y,
	}
}

func FromProto(c *pblogin.LoginReply) *RenWu {
	return &RenWu{
		Name:            c.Name,
		Sex:             c.Sex,
		Life:            c.Life,
		Magic:           c.Magic,
		Physical:        c.Physical,
		Sorcery:         c.Sorcery,
		DefensePhysical: c.DefensePhysical,
		DefenseSorcery:  c.DefenseSorcery,
		X:               c.X,
		Y:               c.Y,
	}
}
