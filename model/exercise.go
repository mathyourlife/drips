package model

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	pb "github.com/mathyourlife/drips/proto"
)

type Exercise struct {
	gorm.Model
	ExerciseClassID int
	ExerciseClass   ExerciseClass
	Modifiers       []Modifier `gorm:"many2many:exercise_modifiers;save_association:false"`
}

func NewExerciseFromProto(e *pb.Exercise) Exercise {
	exercise := Exercise{
		ExerciseClassID: int(e.GetClass().ExerciseClassId),
		ExerciseClass:   NewExerciseClassFromProto(e.GetClass()),
	}

	for _, m := range e.GetModifiers() {
		exercise.Modifiers = append(exercise.Modifiers, NewModifierFromProto(m))
	}
	exercise.ID = uint(e.GetExerciseId())
	return exercise
}

func (m Exercise) String() string {
	s := fmt.Sprintf("%s ", m.ExerciseClass.String())
	var modStrs []string
	for _, mod := range m.Modifiers {
		modStrs = append(modStrs, mod.String())
	}
	if len(modStrs) > 0 {
		s += fmt.Sprintf("(%s) ", strings.Join(modStrs, ", "))
	}
	return strings.TrimSpace(s)
}

func (m Exercise) ToProto() *pb.Exercise {
	ec := &pb.Exercise{
		ExerciseId: int32(m.ID),
		Class:      m.ExerciseClass.ToProto(),
	}
	if len(m.Modifiers) > 0 {
		ec.Modifiers = []*pb.Modifier{}
	}
	for _, mod := range m.Modifiers {
		ec.Modifiers = append(ec.Modifiers, mod.ToProto())
	}
	return ec
}
