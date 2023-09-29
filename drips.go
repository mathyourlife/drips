package drips

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/mathyourlife/drips/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Service struct {
	pb.UnimplementedDripsServiceServer
}

func (s *Service) User(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		User: &pb.User{
			UserId:      2,
			DisplayName: "Dan",
		},
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
						Name: "squat",
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				},
				{
					Sequence: 1,
					Class: &pb.ExerciseClass{
						Name: "squat",
					},
					Duration: &durationpb.Duration{Seconds: 60},
					Rest:     &durationpb.Duration{Seconds: 30},
				},
			},
		},
	}, nil
}

func PrintRoutine(routine *pb.Routine) string {
	var out []string

	for _, e := range routine.Exercises {
		out = append(out, fmt.Sprintf("%d: %s for %d seconds then rest for %d seconds", e.Sequence+1, e.Class.Name, e.Duration.Seconds, e.Rest.Seconds))
	}
	return strings.Join(out, "\n")
}
