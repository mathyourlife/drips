package model

import (
	"gorm.io/gorm"

	pb "github.com/mathyourlife/drips/proto"
)

type Modifier struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

func NewModifierFromProto(mod *pb.Modifier) Modifier {
	m := Modifier{
		Name: mod.GetName(),
	}
	m.ID = uint(mod.GetModifierId())
	return m
}

func (m Modifier) String() string {
	return m.Name
}

func (m Modifier) ToProto() *pb.Modifier {
	return &pb.Modifier{
		ModifierId: int32(m.ID),
		Name:       m.Name,
	}
}
