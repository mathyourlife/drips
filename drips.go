package drips

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	pb "github.com/mathyourlife/drips/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Service struct {
	pb.UnimplementedDripsServiceServer
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) User(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	query := `SELECT user_id, display_name FROM user
		WHERE user_id = $1;`

	u := &pb.User{}
	err := s.db.QueryRowContext(ctx, query, req.UserId).Scan(&u.UserId, &u.DisplayName)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		User: u,
	}, nil
}

func (s *Service) Routine(ctx context.Context, req *pb.RoutineRequest) (*pb.RoutineResponse, error) {
	// Load the routine
	query := `SELECT routine_id, name, source, sequence
		FROM routine
		WHERE routine_id = $1;`

	r := &pb.Routine{}
	err := s.db.QueryRowContext(ctx, query, req.RoutineId).Scan(
		&r.RoutineId, &r.Name, &r.Source, &r.Sequence)
	if err != nil {
		return nil, err
	}

	// Load exercises
	query = `SELECT exercise.exercise_id, exercise.sequence, class.name, exercise.duration_seconds, exercise.rest_seconds
		FROM routine_exercise
		JOIN exercise ON routine_exercise.exercise_id = exercise.exercise_id
		JOIN class ON exercise.class_id = class.class_id
		WHERE routine_exercise.routine_id = $1`

	rows, err := s.db.Query(query, req.RoutineId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		e := &pb.Exercise{
			Class: &pb.ExerciseClass{},
		}
		var durationSec, restSec int
		if err := rows.Scan(&e.ExerciseId, &e.Sequence, &e.Class.Name, &durationSec, &restSec); err != nil {
			return nil, err
		}
		e.Duration = durationpb.New(time.Duration(durationSec) * time.Second)
		e.Rest = durationpb.New(time.Duration(restSec) * time.Second)
		r.Exercises = append(r.Exercises, e)
	}

	// Load modifiers for the exercise
	query = `SELECT modifier.name
		FROM exercise_modifier
		JOIN modifier ON exercise_modifier.modifier_id = modifier.modifier_id
		WHERE exercise_modifier.exercise_id = $1`

	for _, e := range r.Exercises {
		rows, err = s.db.Query(query, e.ExerciseId)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var modifier string
			if err := rows.Scan(&modifier); err != nil {
				return nil, err
			}
			e.Class.Modifiers = append(e.Class.Modifiers, modifier)
		}
	}

	return &pb.RoutineResponse{
		Routine: r,
	}, nil
}

func (s *Service) Routines(ctx context.Context, req *pb.RoutinesRequest) (*pb.RoutinesResponse, error) {

	// Search exercises
	query := `SELECT routine_id, name, source, sequence
		FROM routine;`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var rs []*pb.Routine
	for rows.Next() {
		r := &pb.Routine{}
		if err := rows.Scan(&r.RoutineId, &r.Name, &r.Source, &r.Sequence); err != nil {
			return nil, err
		}
		// Match for a substring of the routine name or match all if search name is empty.
		if strings.Contains(strings.ToLower(r.Name), strings.ToLower(req.Name)) {
			rs = append(rs, r)
		}
	}

	return &pb.RoutinesResponse{
		Routines: rs,
	}, nil
}

func PrintRoutine(routine *pb.Routine) string {
	var exs []string
	for _, e := range routine.Exercises {
		var l string
		l = fmt.Sprintf("%d: %s (%s) for %d seconds", e.Sequence+1, e.Class.Name, strings.Join(e.Class.Modifiers, ","), e.Duration.Seconds)
		if e.Rest != nil && e.Rest.AsDuration() != 0 {
			l += fmt.Sprintf(" then rest for %d seconds", e.Rest.Seconds)
		}

		exs = append(exs, l)
	}

	out := fmt.Sprintf("%s\n#%d\n\n%s", routine.Name, routine.Sequence+1, strings.Join(exs, "\n"))
	return out
}
