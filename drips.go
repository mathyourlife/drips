package drips

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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
	return &pb.RoutineResponse{
		Routine: &pb.Routine{
			Name:     "Caroline Girvan - Iron Series",
			Source:   "https://www.youtube.com/watch?v=SCxNnWW2zB8",
			Sequence: 0, // Day 1
			Exercises: []*pb.Exercise{
				{
					Sequence: 0,
					Class: &pb.ExerciseClass{
						Name:      "squat",
						Modifiers: []string{"suitcase"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 1,
					Class: &pb.ExerciseClass{
						Name:      "squat",
						Modifiers: []string{"suitcase"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 2,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"static", "left"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 3,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"static", "right"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 4,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"static", "left"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 5,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"static", "right"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 6,
					Class: &pb.ExerciseClass{
						Name:      "romanian dead lift",
						ShortName: "rdl",
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 7,
					Class: &pb.ExerciseClass{
						Name:      "romanian dead lift",
						ShortName: "rdl",
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 8,
					Class: &pb.ExerciseClass{
						Name:      "romanian dead lift",
						ShortName: "rdl",
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 9,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"rear step", "left"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 10,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"rear step", "right"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 11,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"rear step", "left"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 12,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"rear step", "right"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 13,
					Class: &pb.ExerciseClass{
						Name:      "squat",
						Modifiers: []string{"goblet", "pause at bottom"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 14,
					Class: &pb.ExerciseClass{
						Name:      "squat",
						Modifiers: []string{"goblet", "pause at bottom"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 15,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"lateral", "left"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 16,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"lateral", "right"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 17,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"lateral", "left"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 18,
					Class: &pb.ExerciseClass{
						Name:      "lunge",
						Modifiers: []string{"lateral", "right"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				}, {
					Sequence: 19,
					Class: &pb.ExerciseClass{
						Name:      "squat",
						Modifiers: []string{"goblet", "1/2 rep"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
				}, {
					Sequence: 20,
					Class: &pb.ExerciseClass{
						Name:      "squat",
						Modifiers: []string{"goblet"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
				}, {
					Sequence: 21,
					Class: &pb.ExerciseClass{
						Name:      "squat",
						Modifiers: []string{"goblet", "1/2 rep"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
				}, {
					Sequence: 22,
					Class: &pb.ExerciseClass{
						Name:      "squat",
						Modifiers: []string{"goblet"},
					},
					Duration: &durationpb.Duration{Seconds: 60},
				},
			},
		},
	}, nil
}

func PrintRoutine(routine *pb.Routine) string {
	var exs []string
	for _, e := range routine.Exercises {
		var l string
		l = fmt.Sprintf("%d: %s (%s) for %d seconds", e.Sequence+1, e.Class.Name, strings.Join(e.Class.Modifiers, ","), e.Duration.Seconds)
		if e.Rest != nil {
			l += fmt.Sprintf(" then rest for %d seconds", e.Rest.Seconds)
		}

		exs = append(exs, l)
	}

	out := fmt.Sprintf("%s\n#%d\n\n%s", routine.Name, routine.Sequence+1, strings.Join(exs, "\n"))
	return out
}
