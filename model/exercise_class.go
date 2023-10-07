package model

import (
	"gorm.io/gorm"

	pb "github.com/mathyourlife/drips/proto"
)

type ExerciseClass struct {
	gorm.Model
	Name      string `gorm:"uniqueIndex"`
	ShortName string
}

func NewExerciseClassFromProto(ec *pb.ExerciseClass) ExerciseClass {
	e := ExerciseClass{
		Name:      ec.GetName(),
		ShortName: ec.GetShortName(),
	}
	e.ID = uint(ec.GetExerciseClassId())
	return e
}

func (m ExerciseClass) String() string {
	return m.Name
}

func (m ExerciseClass) ToProto() *pb.ExerciseClass {
	c := &pb.ExerciseClass{
		ExerciseClassId: int32(m.ID),
		Name:            m.Name,
		ShortName:       m.ShortName,
	}
	return c
}
