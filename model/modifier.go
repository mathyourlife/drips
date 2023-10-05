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
	return Modifier{
		Name: mod.GetName(),
	}
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
