package model

import (
	"fmt"

	"gorm.io/gorm"

	pb "github.com/mathyourlife/drips/proto"
)

type Routine struct {
	gorm.Model
	Name      string
	Source    string
	Sequence  int
	Exercises []Exercise `gorm:"many2many:routine_exercises;"`
}

func NewRoutineFromProto(r *pb.Routine) Routine {
	routine := Routine{
		Name:     r.GetName(),
		Source:   r.GetSource(),
		Sequence: int(r.GetSequence()),
	}

	for _, e := range r.GetExercises() {
		routine.Exercises = append(routine.Exercises, NewExerciseFromProto(e))
	}

	return routine
}

func (m Routine) String() string {
	return fmt.Sprintf("%s %d", m.Name, m.Sequence)
}

func (m Routine) ToProto() *pb.Routine {
	r := &pb.Routine{
		RoutineId: int32(m.ID),
		Name:      m.Name,
		Source:    m.Source,
		Sequence:  int32(m.Sequence),
	}
	if len(m.Exercises) > 0 {
		r.Exercises = []*pb.Exercise{}
	}
	for _, e := range m.Exercises {
		r.Exercises = append(r.Exercises, e.ToProto())
	}
	return r
}
