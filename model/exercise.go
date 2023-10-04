package model

import (
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/gorm"

	pb "github.com/mathyourlife/drips/proto"
)

type Exercise struct {
	gorm.Model
	Sequence  int
	Class     ExerciseClass `gorm:"embedded"`
	Modifiers []Modifier    `gorm:"many2many:exercise_modifiers;"`
	Duration  time.Duration
	Rest      time.Duration
	Reps      int
}

func NewExerciseFromProto(e *pb.Exercise) Exercise {
	exercise := Exercise{
		Sequence: int(e.GetSequence()),
		Class:    NewExerciseClassFromProto(e.GetClass()),
		Duration: e.GetDuration().AsDuration(),
		Rest:     e.GetRest().AsDuration(),
		Reps:     int(e.GetReps()),
	}

	for _, m := range e.GetModifiers() {
		exercise.Modifiers = append(exercise.Modifiers, NewModifierFromProto(m))
	}

	return exercise
}

func (m Exercise) String() string {
	s := fmt.Sprintf("%s ", m.Class.String())
	var modStrs []string
	for _, mod := range m.Modifiers {
		modStrs = append(modStrs, mod.String())
	}
	if len(modStrs) > 0 {
		s += fmt.Sprintf("(%s) ", strings.Join(modStrs, ", "))
	}
	if m.Duration > 0 {
		s += fmt.Sprintf("for %s ", m.Duration)
	}
	if m.Reps > 0 {
		s += fmt.Sprintf("for %d reps ", m.Reps)
	}
	if m.Rest > 0 {
		s += fmt.Sprintf("and then rest for %s", m.Rest)
	}
	return strings.TrimSpace(s)
}

func (m Exercise) ToProto() *pb.Exercise {
	ec := &pb.Exercise{
		ExerciseId: int32(m.ID),
		Sequence:   int32(m.Sequence),
		Class:      m.Class.ToProto(),
		Duration:   durationpb.New(m.Duration),
		Rest:       durationpb.New(m.Rest),
		Reps:       int32(m.Reps),
	}
	if len(m.Modifiers) > 0 {
		ec.Modifiers = []*pb.Modifier{}
	}
	for _, mod := range m.Modifiers {
		ec.Modifiers = append(ec.Modifiers, mod.ToProto())
	}
	return ec
}
